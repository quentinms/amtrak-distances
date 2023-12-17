# Amtrak Distances

Distances along Amtrak routes

For my [Coast Starlight](https://blog.quentin.ms/posts/coast-starlight/) blog post, I needed the distance by train between Emeryville and Seattle and for some reason search engines cannot find anything useful (and LLMs just hallucinate).

It's not on Amtrak's website either. They used to have them as part of the timetables, but they're not on the newer ones.

I had to resort to finding what I was looking for in [this unofficial archive of Amtrak timetables](https://juckins.net/amtrak_timetables/archive/home.php).

Results are visible [here](https://quentin.ms/amtrak-distances/).


## Development

It uses Go templates to generate a single HTML page based on `routes.json`.

```sh
go run main.go | tee index.html
```
