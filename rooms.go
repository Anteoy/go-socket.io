package socketio

var Roomsmap map[string][]*NamespaceConn = make(map[string][]*NamespaceConn, 0)
