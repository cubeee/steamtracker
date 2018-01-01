{% import "_macros.tpl" game_image, time_played %}
<table class="ui celled table game-table main-game-table">
    <tbody>
    {% for game_stat in game_stats %}
        <tr>
            <td>
                <h4 class="ui image header">
                    {{ game_image(game_stat.Game.AppId, game_stat.Game.Icon, "ui mini rounded image") }}
                    <div class="content">
                        <a href="/game/{{ game_stat.Game.AppId }}/">{{ game_stat.Game.Name }}</a>
                        <div class="sub header" title="{{ game_stat.MinutesPlayed }} minutes">
                            {{ time_played(game_stat.MinutesPlayed) }}
                        </div>
                    </div>
                </h4>
            </td>
        </tr>
    {% endfor %}
    {% if fill_table && game_stats|length < max_rows %}
        {% loop max_rows-game_stats|length %}
        <tr>
            {% loop columns %}
            <td class="empty">&nbsp;</td>
            {% endloop %}
        </tr>
        {% endloop %}
    {% endif %}
    </tbody>
</table>