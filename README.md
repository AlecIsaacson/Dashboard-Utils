# Dashboard-Exporter
This app will export all dashboards from a New Relic account in the standard JSON format, storing each dashboard definition in
a separate file named after the dashboard ID, which is an integer.

It expects one command line argument:
  -apikey (the API key for the account the dashboard is to be exported from)

Once you have these files, you can store them as backups, or edit them for importing into another account using the
dashboard-importer.



# Dashboard-Importer
This app will take a New Relic dashboard JSON definition and import it into New Relic Insights using the Dashboards API.
It expects two command line arguments:
  -apikey (the API key for the account the dashboard is to be imported into)
  -defn (the path to and file name of the dashboard definition JSON file)

This app does not do any error checking.  If you are transferring dashboards between accounts, be sure to remove all references to account_id and drilldown_dashboard_id.

See https://newrelic.jiveon.com/people/aisaacson@newrelic.com/blog/2018/03/05/exporting-and-importing-insights-dashboards for more details and some helpful tips.



# Dashboard-Cleaner
This app takes a New Relic dashboard JSON definition file and strips out the account info, URLs, and links so it can be imported into another account without leaking customer information or showing unexpected data.
It expects one command line argument:
  -file (the name of the dashboard file to clean)
  
The cleaner outputs a new file named after the dashboard's title.
