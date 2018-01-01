package tags

import (
	"github.com/flosch/pongo2"
)

type tagLoopNode struct {
	startEvaluator pongo2.IEvaluator
	countEvaluator pongo2.IEvaluator
	reversed       bool

	wrapper *pongo2.NodeWrapper
}

type tagLoopInformation struct {
	Counter    int
	First      bool
	Last       bool
	ParentLoop *tagLoopInformation
}

func (node *tagLoopNode) Execute(ctx *pongo2.ExecutionContext, writer pongo2.TemplateWriter) (loopError *pongo2.Error) {
	loopCtx := pongo2.NewChildExecutionContext(ctx)
	parentLoop := loopCtx.Private["loop"]

	// Create loop struct
	loopInfo := &tagLoopInformation{
		First: true,
	}

	// Is it a loop in a loop?
	if parentLoop != nil {
		loopInfo.ParentLoop = parentLoop.(*tagLoopInformation)
	}

	// Register loopInfo in public context
	loopCtx.Private["forloop"] = loopInfo

	start := 1
	count := 1

	if node.startEvaluator != nil {
		obj, err := node.startEvaluator.Evaluate(loopCtx)
		if err != nil {
			return err
		}
		if !obj.IsInteger() {
			return loopCtx.Error("Expression has to resolve to an integer", nil)
		}
		start = obj.Integer()
	}

	if node.countEvaluator != nil {
		obj, err := node.countEvaluator.Evaluate(loopCtx)
		if err != nil {
			return err
		}
		if obj.IsNil() {
			count = start
			start = 0
		} else if obj.IsInteger() {
			count = obj.Integer()
		} else {
			return loopCtx.Error("Expression has to resolve to an integer", nil)
		}
	}

	if node.reversed {
		count = -count
	}

	inLoop := func(idx int, first bool, last bool) bool {
		loopInfo.Counter = idx
		loopInfo.First = first
		loopInfo.Last = last

		// Render elements with updated context
		err := node.wrapper.Execute(loopCtx, writer)
		if err != nil {
			loopError = err
			return false
		}
		return true
	}

	end := start + count
	if count < 0 { // reverse
		for i := start; i > end; i-- {
			inLoop(i, i == start, i == end)
		}
	} else if count >= 0 { // normal
		for i := start; i < end; i++ {
			inLoop(i, i == start, i == end)
		}
	}
	return loopError
}

func LoopTagParser(doc *pongo2.Parser, _ *pongo2.Token, arguments *pongo2.Parser) (pongo2.INodeTag, *pongo2.Error) {
	node := &tagLoopNode{}

	wrapper, endArgs, err := doc.WrapUntilTag("endloop")
	if err != nil {
		return nil, err
	}
	node.wrapper = wrapper

	if endArgs.Count() > 0 {
		return nil, endArgs.Error("Arguments not allowed here", nil)
	}

	firstEvaluator, err := arguments.ParseExpression()
	if err != nil {
		return nil, err
	}
	arguments.Consume()
	node.countEvaluator = firstEvaluator

	secondEvaluator, err := arguments.ParseExpression()
	if err == nil {
		node.startEvaluator = firstEvaluator
		node.countEvaluator = secondEvaluator
	}
	arguments.Consume()

	if arguments.MatchOne(pongo2.TokenIdentifier, "reversed") != nil {
		node.reversed = true
	}
	arguments.Consume()

	return node, nil
}
