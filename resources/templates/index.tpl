{% extends "_base.tpl" %}

{% block navigation %}{% endblock %}

{% block body %}
{% import "_macros.tpl" time_played %}
<div class="pusher">
    <div class="ui inverted vertical masthead center aligned segment">
        <div class="ui text container">
            <h1 class="ui inverted header">SteamTracker</h1>
            <h2>Track your Steam profile progression</h2>
            <div class="ui ten column centered tight grid">
                <div class="row">
                    <div class="ui twelve wide column">
                        <form class="ui form" method="post" action="/search">
                            <div class="fields">
                                <div class="sixteen wide centered field">
                                    <div class="ui fluid big icon input">
                                        <!-- TODO: csrf -->
                                        <input class="prompt" name="identifier" type="text" placeholder="Steam profile url, id or custom name">
                                        <i class="search icon"></i>
                                    </div>
                                </div>
                            </div>
                        </form>
                    </div>
                </div>
                <div class="row">
                    <div class="ui twelve wide center aligned column">
                        <a class="ui large black image label">
                            {{ tracked_players|formatnumber:"%d" }}
                            <div class="detail">Players tracked</div>
                        </a>
                        <a class="ui large black image label">
                            {{ time_played(collective_hours_tracked, false) }}
                            <div class="detail">Collective hours tracked</div>
                        </a>
                    </div>
                </div>
            </div>
        </div>
    </div>

    <div class="main ui relaxed padded stackable centered grid">
        <div class="four wide column">
            <h3 class="text-medium">Most played in the last 24 hours</h3>
            {% include "components/game_table.tpl" with fill_table=true max_rows=10 columns=1 game_stats=game_stats_24h %}
        </div>
        <div class="four wide column">
            <h3 class="text-medium">Most played in the last 7 days</h3>
            {% include "components/game_table.tpl" with fill_table=true max_rows=10 columns=1 game_stats=game_stats_7d %}
            </div>
        <div class="four wide column">
            <h3 class="text-medium">Most played</h3>
            {% include "components/game_table.tpl" with fill_table=true max_rows=10 columns=1 game_stats=game_stats %}
        </div>
    </div>
</div>
{% endblock %}