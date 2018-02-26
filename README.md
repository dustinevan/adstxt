# Crawler
This ads.txt crawler, is written to the ads.txt specification found here https://iabtechlab.com/wp-content/uploads/2017/09/IABOpenRTB_Ads.txt_Public_Spec_V1-0-1.pdf. It handles redirects, correctly deals with subdomains, and handles common publisher implementation mistakes. The populated adstxt file object contains not only information about the ad paths, but also information about erroneous rows, inline comments, variables, etc. 

The file object also contains information about the crawl of the ads.txt, including the root domain, the domain for which the ads.txt is valid, the url of the ads.txt, The lookup time, and a file checksum. 

## Finding the correct ads.txt
There is some confusion about which paths to put in an ads.txt file and where that file should be located. For example many blogging domains serve the global ads.txt file on each blog subdomain. The crawler traverses the url subdomains to figure out which domain the ads.txt is valid for. Given ```myblog.blogdomain.com``` the crawler will ask for ```myblog.blogdomain.com/ads.txt```. If an ads.txt is found the crawler will then request ```blogdomain.com/ads.txt``` and compare checksums to see if the files are the same. This allows users of the crawler to avoid the possible n^2 duplication of ads.txt paths on blog domains.

## Example Usage:
```git clone https://github.com/dustinevan/adstxt.git```
```cd ./adstxt/app/atcrawl```
```./atcrawl bloomberg.com/ads.txt```
```./atcrawl batman.wikia.com/ads.txt```

The example binary was built on masOS Sierra 10.12.6 with go1.9.2 darwin/amd64

