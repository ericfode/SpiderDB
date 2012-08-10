SpiderDB
========

Graph database backed by Redis.

SpiderDB uses Redis as a backing datastore.

TODO: Create Response to hash function for getting nodes
TODO: Create function for converting byte arrays to ints more quickly

First Phase : Structure
Nodes in spiderDB are stored as hashes with each proprty as a key in the hash
Edges are stored as a compositie id of the two nodes they connect (2to3)
Adjacencys are stored in a sorted set using rank as one of the ids
Add commands Cache parts of graph in application memeory

Second Phase : Genralization
Add ablitiy for user to create search and indexing functions
Set up go project so that the backing store can be changed (to foundationDB, or rethinkDB possibly)

Third Phase : Querying





SpidySource (the first use)
Will be used as a version control system that knows your code
