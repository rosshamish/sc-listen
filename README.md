```
sc-listen master ❯ sc-listen help
NAME:
   sc-listen - stream soundcloud from the terminal

USAGE:
   sc-listen [global options] command [command options] [arguments...]
   
VERSION:
   0.0.0
   
COMMANDS:
   help, h	Shows a list of commands or help for one command
   
GLOBAL OPTIONS:
   --help, -h		show help
   --version, -v	print the version
```

Build

`go build`

Run

`./sc-listen [username]`

```
sc-listen master* ❯ ./sc-listen troyboi
Did you mean...
	%s TroyBoi
	%s TroyBoi Onlygodcanjudgeme
	%s troyboi924

Found user TroyBoi
	Full Name: 
	Country: 
Tracks:
	TroyBoi: TroyBoi Feat NEFERA - On My Own
	TroyBoi: Tinie Tempah feat Jess Glynne - Not Letting Go (TroyBoi Remix)
	TroyBoi: Wonky
	TroyBoi: TroyBoi - BBC Radio 1 Guest Mix - DJ Target
	TroyBoi: ili
	...
Playlists:
	TroyBoi: TrapMusic.NET Uploads
	TroyBoi: SoundSnobz
	TroyBoi: icekream
```

---

useful resources:
- github.com/moriyoshi/pulsego/ to play to speakers
- golang-soundcloud as a soundcloud api wrapper (and probably contribute back to this)
- github.com/mdlayher/waveform for printing a waveform soundcloud-style
- https://developers.soundcloud.com/docs/api/reference

