package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"net/http"
	"os"

	"github.com/tune-bot/core"
	"github.com/valyala/fastjson"
)

func Add(w http.ResponseWriter, req *http.Request) {
	var parser fastjson.Parser
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(req.Body)
	raw, err := parser.Parse(buf.String())

	if err != nil {
		errorResponse(err, w)
		return
	}

	song := core.Song{
		Code:   string(raw.GetStringBytes("code")),
		Title:  string(raw.GetStringBytes("title")),
		Artist: string(raw.GetStringBytes("artist")),
		Album:  string(raw.GetStringBytes("album")),
		Year:   uint16(raw.GetInt("year")),
	}

	err = song.AddToPlaylist(string(raw.GetStringBytes("playlistId")))
	if err != nil {
		errorResponse(err, w)
		return
	}
	data, err := json.Marshal(song)

	successResponse(data, w)
}

func Remove(w http.ResponseWriter, req *http.Request) {
	var parser fastjson.Parser
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(req.Body)
	raw, err := parser.Parse(buf.String())

	if err != nil {
		errorResponse(err, w)
		return
	}

	song := core.Song{Id: string(raw.GetStringBytes("songId"))}

	err = song.RemoveFromPlaylist(string(raw.GetStringBytes("playlistId")))
	if err != nil {
		errorResponse(err, w)
		return
	}

	successResponse([]byte("{}"), w)
}

func Download(w http.ResponseWriter, req *http.Request) {
	song := core.Song{}
	err := json.NewDecoder(req.Body).Decode(&song)

	if err != nil {
		errorResponse(err, w)
		return
	}

	file, err := os.Open(song.FilePath())
	if err != nil {
		errorResponse(err, w)
		return
	}
	defer file.Close()

	stats, err := file.Stat()
	if err != nil {
		errorResponse(err, w)
		return
	}

	var size int64 = stats.Size()
	bytes := make([]byte, size)

	bufr := bufio.NewReader(file)
	_, err = bufr.Read(bytes)

	if err != nil {
		errorResponse(err, w)
		return
	}

	w.Header().Set("Content-Type", "audio/mpeg")
	w.Header().Set("Content-Disposition", "filename="+file.Name())
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)

	return
}

func Search(w http.ResponseWriter, req *http.Request) {
	query := "Drok Drok Drok"

	results := core.Search(query) // TODO: numResults

	data, err := json.Marshal(results)

	if err != nil {
		errorResponse(err, w)
		return
	}

	successResponse(data, w)
}
