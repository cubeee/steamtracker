{% import "_macros.tpl" google_analytics %}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8" />
    <meta http-equiv="X-UA-COMPATIBLE" content="IE=Edge" />
    <meta name="application-name" content="SteamTracker" />
    <meta name="author" content="cubeee" />
    <meta name="rating" content="general" />
    <meta name="keywords" content="Steam, playtime, tracker, tracking" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0, user-scalable=yes" />
    <meta name="apple-mobile-web-app-capable" content="yes" />
    <meta name="apple-mobile-web-app-status-bar-style" content="blue-translucent" />
    <meta name="theme-color" content="#2c3e50" />
    <meta property="og:title" content="Steam Tracker" />
    <meta property="og:site_name" content="SteamTracker" />
    <meta property="og:description" content="Steam user tracking site" />
    <meta property="og:type" content="website" />

    <title>{% if pageTitle %}{{ pageTitle }} - SteamTracker{% else %}SteamTracker{% endif %}</title>
    {% block stylesheets %}
    <link rel="stylesheet" type="text/css" href="/static/global.css" />
    {% endblock %}
    <script type="application/javascript" src="/static/global.js"></script>
    {% block header_scripts %}

    {% endblock %}
</head>
<body>
    {% block navigation %}
    <div class="ui inverted vertical segment">
        <div class="ui grid">
            <div class="twelve wide centered column" style="padding-top: 0; padding-bottom: 0;">
                <div class="ui huge inverted menu">
                    <div class="header item"><a href="/">SteamTracker</a></div>
                    <div class="right menu">
                        <div class="item">
                            <form class="ui icon input" method="post" action="/search">
                                <div class="ui transparent inverted icon input">
                                    <i class="search icon"></i>
                                    <input type="hidden" name="${_csrf.parameterName}" value="${_csrf.token}" />
                                    <input type="text" name="identifier" placeholder="Search player">
                                </div>
                            </form>
                        </div>
                        <!-- TODO: authenticated -->
                    </div>
                </div>
            </div>
        </div>
    </div>
    {% endblock %}

    {% block body %}{% endblock %}

    {% block footer %}
    <div class="ui inverted vertical masthead center aligned segment">
        <!-- TODO: fill -->
        <p>SteamTracker</p>
    </div>
    {% endblock %}
    {% block footer_scripts %}{% endblock %}
    {% block analytics %}
        {{ google_analytics(googleAnalyticsId) }}
    {% endblock %}
</body>
</html>
