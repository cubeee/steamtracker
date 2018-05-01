{% macro google_analytics(id) export %}
    {% if id %}
        <script>
            !function(t,r,a,c,k,e,d){t.GoogleAnalyticsObject=a;t[a]||(t[a]=function(){
                (t[a].q=t[a].q||[]).push(arguments)});t[a].l=+new Date;e=r.createElement(c);
                d=r.getElementsByTagName(c)[0];e.src=k;d.parentNode.insertBefore(e,d)}
            (window,document,'ga','script','//www.google-analytics.com/analytics.js');
            ga('create', '{{ id }}', 'auto');
            ga('send', 'pageview');
        </script>
    {% endif %}
{% endmacro %}

{% macro game_image(id, url, classes) export %}
<img src="https://steamcdn-a.akamaihd.net/steamcommunity/public/images/apps/{{ id }}/{{ url }}.jpg" class="{{ classes }}" />
{% endmacro %}

{% macro time_played(minutes, detailed=true) export %}
    {% set hours=minutes/60|integer %}
    {% if detailed %}
        {% if minutes <= 60 %}
            {{ formatted_number(minutes) }} minute{{ minutes|pluralize }}
        {% else %}
            {{ formatted_number(hours) }} hour{{ hours|pluralize }}
        {% endif %}
    {% else %}
        {{ formatted_number(hours) }}
    {% endif %}
{% endmacro %}

{% macro formatted_number(num) export %}{{ num|formatnumber:"%d" }}{% endmacro %}