# poe-ML-indexer
App that reads the Path of Exile item river API and saves all identified rares that have a buyout in a mongo database.

You have to set your own database as an environment variable before launching the app.
Ex : SET db=mongodb://test:test@mymongodb.com:1337/DBNAME

Using a local mongo DB (localhost) is advised
