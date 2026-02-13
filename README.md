# Opal Media Server

***Opal is currently pre-alpha software. It is not intended for daily-driving yet and lacks many features including lacking support for video streaming entirely. It also has very poor documentation and instructions currently. Progress towards an alpha release is rapidly developing so check back in soon!***

Opal is a personal media management server written in golang. It implements the Jellyfin API allowing for the use of any client in the existing Jellyfin ecosystem like jellyfin-web or the jellyfin android client.

**Opal is intentionally very opinionated. Opal is intended for admins who would rather use ssh and config files to maintain their server than a web dashboard. Opal does not implement the API endpoints required for the use of the jellyfin admin dashboard. Opal is unlikely to ever implement Live TV, Music or Book support. Instead the developer recommends navidrome for personal management of music**

## Features

- **Memory usage:** Opal internally is much simpler than jellyfin and golang's runtime is much smaller than .Net's runtime. When idle with one client connected on a test machine, jellyfin uses ~130MiB of ram and opal uses ~13MiB (Note: idle memory usage will increase as more features are added).
- **No metadata sidecaring:** Jellyfin stores metadata for your personal media inside the directory containing the media. Opal stores its metadata for your personal media in a configurable location seperate from your personal media collection. This means the user running the opal media server binary only needs read and execute permissions to serve your personal media collections.
- **Hardened Authentication (argon2id):** Opal uses argon2id with very secure parameters by default. This causes much slower login times compared to jellyfin but provides industry-level resistance to GPU-accelerated and classical bruteforce attacks.
- **Forced media naming scheme:** Opal only recognises media if it's filename has a substring that matches following pattern: <br>`[{tvshow or movie}-imdb-{imdb id}]` Examples: `[movie-imdb-tt0078748]`, `[tvshow-imdb-tt0412142]`<br>
Movies must be a file (ie not folder with the primary mkv or mp4 inside) and tv shows must be folders which contain seasons (folders) that have identifying substrings in the name (Examples: `S01[s1]`, `Season 2[s2]`, `S3[s3]` etc.) and episode files with substring filenames inside those season folders (Examples: `episode 1[ep1].mkv`, `Episode 02[ep2].mp4`, `[ep3]Episode 3.mkv` etc.). This restrictive naming scheme allows for extremely fast server initialisation, metadata retrieval, and media serving.

## Installation and Setup

*Opal is pre-alpha software, expect bugs, crashes, incomplete features, and incompatibilities with some jellyfin clients.*


1. Modify default configs:

/config/server.cfg:
```
address=0.0.0.0
port=7096
metadata_dir="./metadata"
cache_dir="./cache"
```

2. Create library configs

/config/libraries/example_lib.cfg:
```
display_name="Display name"
path="/path/to/library"
```
*Note: To create multiple libraries, just add another cfg file to config/libraries and restart the server*

3. Create a TMDB config file:

/config/tmdb.cfg
```
# Opal uses TMDB read-access tokens (v4), not v3 api keys
# Direct storage:
api_key="insert TMDB read-acess token here"

# Or via environment variable:
api_key=$TMDB_TOKEN
```

4. Install a jellyfin web frontend:

Download a release build of jellyfin (10.11.6 preferably) and extract `jellyfin-web` to the root of the repository folder. Then rename the folder from `jellyfin-web` to `web`.

5. Create a user
```
go run ./cmd/opal-new-user
```

6. Build or run the server:
```
go run ./cmd/opal
```
or
```
go build ./cmd/opal && ./opal
```

## Contributions

Contributions are welcome. If you are looking for easy ways to help, have a look in `/internal/httpserver/api`. Most API endpoints are stubbed at least partially but wouldn't be terribly hard to complete. Also check out `/todo.txt`. However I ask that you refrain from writing major features, Opal is designed to have a very limited scope and all major features are already implemented or are a work in process. Adding support for the jellyfin admin panel, other media types, or general feature creeping will not be merged. 

## License

Distributed under the GPLv3 License. See LICENSE for more information.
