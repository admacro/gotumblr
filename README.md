# gotumblr
CLI client for posting text and multiple quotes to Tumblr.

## Requirement
 - Get your Tumblr API key at https://www.tumblr.com/oauth/apps
 - Get your Tumblr oauth token at https://api.tumblr.com/console/calls/user/info

## How to use it
First, put your secret keys and oauth credentials in your environment.
```bash
export TUMBLR_BLOG_NAME=your-blog-name
export TUMBLR_CONSUMER_KEY=aaa
export TUMBLR_CONSUMER_SECRET=bbb
export TUMBLR_OAUTH_TOKEN=ccc
export TUMBLR_OAUTH_TOKEN_SECRET=ddd
```

### Post multiple quotes
1. Put quotes and source into `quotes.txt` like following. The first
line is the source of the quotes, the following the quotes.
```txt
中国古人
学而不思则惘，思而不学则怠。
俭以养德，静以修身。
净口、修身、齐家、治国、平天下。
```
2. run `./gotumblr q`

### Post text
1. Put title and text into `text.md` like following. The first
line is the title, the following the body of the text. The format of
the text is set to `markdown`.
```markdown
《孟子·告子下》
故天将降大任于斯人也，必先苦其心志，劳其筋骨，饿其体肤，空乏其身，行拂乱其所为，所以动心忍性，曾益其所不能。
```
2. run `./gotumblr t`

### Ruby version
There is a similar version written in Ruby [here](https://github.com/admacro/tumblr-batch-quotes).
