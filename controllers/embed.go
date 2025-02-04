package controllers

import (
	"net/http"
)

// GetChatEmbedreadwrite gets the embed for readwrite chat.
func GetChatEmbedreadwrite(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "webroot/index-standalone-chat-readwrite.html")
}

// GetChatEmbedreadonly gets the embed for readonly chat.
func GetChatEmbedreadonly(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "webroot/index-standalone-chat-readonly.html")
}

// GetVideoEmbed gets the embed for video.
func GetVideoEmbed(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "webroot/index-video-only.html")
}
