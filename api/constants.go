package main

var version string

// omega service endpoints to update last mods for pages

const StagingEndPoint string = "https://gw.getmega.dev/omega/twirp/omega.pb.Website/UpdateUrlLastModTime"
const ProdEndPoint string = "https://gw.getmega.co/omega/twirp/omega.pb.Website/UpdateUrlLastModTime"

const CommentoProdDomain string = "www.getmega.com"
const CommentoStagingDomain string = "getmega-app.web.app"
