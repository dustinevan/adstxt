package adstxt



var exdoms = map[string]string{"rubicon.com": "Rubicon Project",
	"fastlane.rubiconproject.com":                   "Rubicon Project",
	"ads.rubiconproject.com":                        "Rubicon Project",
	"rubiconproject.com":                            "Rubicon Project",
	"rubiconproject.com<http://rubiconproject.com>": "Rubicon Project",
	"33across.com":                                  "33Across",
	"pubmatic.com":                                  "PubMatic",
	"apps.pubmatic.com":                             "PubMatic",
	"pubmatic":                                      "PubMatic",
	"openx.com":                                     "OpenX",
	"openx":                                         "OpenX",
	"openxebda":                                     "OpenX",
	"openxprebid":                                   "OpenX",
	"openx.com<http://openx.com>":                   "OpenX",
	"openx.net":                                     "OpenX",
	"facebook.com":                                  "Facebook",
	"facebook":                                      "Facebook",
	"facebook:facebook.com":                         "Facebook",
	"gumgum.com":                                    "GumGum",
	"kargo.com":                                     "Kargo",
	"google.com":                                    "Google",
	"googletagservices.com":                         "Google",
	"?google.com":                                   "Google",
	"adsense":                                       "Google",
	"google.com/adsense":                            "Google",
	"google.com<http://google.com>":                 "Google",
	"www.google.com/dfp":                            "Google",
	"brealtime.com":                                 "bRealtime",
	"Brealtime":                                     "bRealtime",
	"brealtimegoogle":                               "bRealtime",
	"emxdgt.com105":                                 "bRealtime",
	"amazon-adsystem.com":                           "Amazon",
	"c.amazon-adsystem.com":                         "Amazon",
	"advertising.amazon.com":                        "Amazon",
	"amazon.com":                                    "Amazon",
	"a9.com":                                        "Amazon",
	"aps.amazon.com":                                "Amazon",
	"adtech.com":                                    "One by AOL: Display",
	"adtech.net":                                    "One by AOL: Display",
	"aolcloud.net":                                  "One by AOL: Display",
	"liveintent.com":                                "LiveIntent",
	"yieldmo.com":                                   "Yieldmo",
	"mopub.com":                                     "MoPub",
	"aol.com":                                       "One by AOL: Mobile",
	"smartstream.tv":                                "SmartStream",
	"smaato.com":                                    "Smaato",
	"spx.smaato.com":                                "Smaato",
	"taboola.com":                                   "Taboola",
	"trustx.org":                                    "TrustX",
	"sofia.trustx.org":                              "TrustX",
	"lkqd.net":                                      "LKQD",
	"lkqd.com":                                      "LKQD",
	"ad.lkqd.net":                                   "LKQD",
	"criteo.com":                                    "Criteo",
	"critero.com":                                   "Criteo",
	"criteo.net":                                    "Criteo",
	"phillymag.com==criteo.com":                     "Criteo",
	"exponential.com":                               "Exponential",
	"exponential.comi":                              "Exponential",
	"xponential.com":                                "Exponential",
	"lijit.com":                                     "Sovrn",
	"meridian.sovrn.com":                            "Sovrn",
	"sovrn.com":                                     "Sovrn",
	"lijit":                                         "Sovrn",
	"rhythmone.com":                                 "RhythmOne",
	"1rx.io":                                        "RhythmOne",
	"yldbt.com":                                     "Yieldbot",
	"technorati.com":                                "Technorati",
	"bidfluence.com":                                "Bidfluence",
	"beachfront.com":                                "Bidfluence",
	"switch.com":                                    "Switch Concepts",
	"switchconcept":                                 "Switch Concepts",
	"switchconcepts.com":                            "Switch Concepts",
	"brightroll.com":                                "BrightRoll from Yahoo!",
	"conversantmedia.com":                           "Conversant",
	"go.sonobi.com":                                 "Sonobi",
	"sonobi.com":                                    "Sonobi",
	"*.go.sonobi.com":                               "Sonobi",
	"spoutable.com":                                 "Spoutable",
	"freewheel.tv":                                  "FreeWheel",
	"cdn.stickyadstv.com":                           "FreeWheel",
	"stickyad:freewheel.tv":                         "FreeWheel",
	"connatix.com":                                  "Connatix",
	"t.brand-server.com":                            "Centro Brand Exchange",
	"positivemobile.com":                            "Positive Mobile",
	"memeglobal.com":                                "MemeGlobal",
	"kixer.com":                                     "Kixer",
	"sekindo.com":                                   "Sekindo",
	"sekindo":                                       "Sekindo",
	"360yield.com":                                  "Improve Digital",
	"improvedigital.com":                            "Improve Digital",
	"adform.com":                                    "AdForm",
	"adform.net":                                    "AdForm",
	"adx.adform.net":                                "AdForm",
	"inner-active.com":                              "Inneractive",
	"spotxchange.com":                               "SpotX",
	"spotx.tv":                                      "SpotX",
	"streamrail.net":                                "StreamRail",
	"sdk.streamrail.com":                            "StreamRail",
	"mathtag.com":                                   "MediaMath",
	"mediamath.com":                                 "MediaMath",
	"adyoulike.com":                                 "AdYouLike",
	"indexexchnage.com":                             "Index Exchange",
	"indexexchange.com":                             "Index Exchange",
	"www.indexexchange.com":                         "Index Exchange",
	"indexechange.com":                              "Index Exchange",
	"indexexchange(ebda)":                           "Index Exchange",
	"indexexchange(pubmatic)":                       "Index Exchange",
	"indexexchange(videossp)":                       "Index Exchange",
	"index.com":                                     "Index Exchange",
	"kiosked.com":                                   "Kiosked",
	"ads.kiosked.com":                               "Kiosked",
	"video.unrulymedia.com":                         "UnrulyX",
	"brightcom.com":                                 "Brightcom",
	"rs-stripe.com":                                 "PowerInbox",
	"fyber.com":                                     "Fyber",
	"tidaltv.com":                                   "TidalTV",
	"nativo.com":                                    "Nativo",
	"jadserve.postrelease.com":                      "Nativo",
	"media.net":                                     "Media.net",
	"www.yumenetworks.com":                          "YuMe",
	"yume.com":                                      "YuMe",
	"yumenetworks.com":                              "YuMe",
	"revcontent.com":                                "RevContent",
	"revontent.com":                                 "RevContent",
	"outbrain.com":                                  "Outbrain",
	"zedo.com":                                      "Zedo",
	"freeskreen.com":                                "SlimCut Media",
	"bidtellect.com":                                "Bidtellect",
	"smartadserver.com":                             "Smart RTB+",
	"loopme.com":                                    "LoopMe",
	"vidazoo.com":                                   "Vidazoo",
	"vidazoo":                                       "Vidazoo",
	"videoflare.com":                                "Videoflare",
	"yahoo.com":                                     "Gemini from Yahoo!",
	"pixfuture.com":                                 "PixFuture",
	"oms.eu":                                        "OMS",
	"stroeer.com":                                   "Ströer",
	"emxdgt.com":                                    "bRealtime",
	"c1exchange.com":                                "C1X",
	"synacor.com":                                   "Synacor",
	"sfx.freewheel.tv":                              "FreeWheel",
	"videologygroup.com":                            "Videology",
	"tremorhub.com":                                 "Telaria (fka Tremor Video)",
	"altitudedigital.com":                           "Altitude Digital",
	"platform.videologygroup.com":                   "Videology",
	"imonomy.com":                                   "Imonomy",
	"komoona ltd":                                   "Komoona",
	"komoonaltd":                                    "Komoona",
	"springserve.com":                               "SpringServe",
	"spingserve.com":                                "SpringServe",
	"triplelift.com":                                "TripleLift",
	"www.triplelift.com":                            "TripleLift",
	"ib.adnxs.com":                                  "AppNexus",
	"appnexus.com":                                  "AppNexus",
	"appnexus":                                      "AppNexus",
	"apnexus.com":                                   "AppNexus",
	"appnexus.txt":                                  "AppNexus",
	"adnxs.com":                                     "AppNexus",
	"appnexus.com<http://appnexus.com>":             "AppNexus",
	"s.ntv.io/serve":                                "NTV",
	"coxmt.com":                                     "COMET",
	"undertone.com":                                 "Undertone",
	"advertising.com":                               "One by AOL: Video",
	"c.algovid.com":                                 "Algovid",
	"lockerdome.com":                                "Lockerdome",
	"widespace.com":                                 "Widespace",
	"deployads.com":                                 "Sortable",
	"www.mobfox.com":                                "Mobfox",
	"mobfox.com":                                    "Mobfox",
	"teads.tv":                                      "Teads",
	"teads.com":                                     "Teads",
	"publishers.teads.tv":                           "Teads",
	"contextweb.com":                                "PulsePoint",
	"pulsepoint.com":                                "PulsePoint",
	"pulsepoint":                                    "PulsePoint",
	"pulsepoint:contextweb.com":                     "PulsePoint",
	"districtm.com":                                 "District M",
	"districtm.ca":                                  "District M",
	"districtm.io":                                  "District M",
	"sharethrough.com":                              "Sharethrough",
	"media.adfrontiers.com":                         "Adfrontiers",
	"adfrontiers.com":                               "Adfrontiers",
	"media.adfrontiers":                             "Adfrontiers",
	"ad3media.com":                                  "Ad3media",
	"ads.admized.com":                               "ADMIZED",
	"admized.com":                                   "ADMIZED",
	"a.twiago.com":                                  "Twiago",
	"twiago.com":                                    "Twiago",
	"xapads.com":                                    "Xapads",
	"ad-stir.com":                                   "Adstir",
	"ad.yieldlab.net":                               "Yieldlab",
	"yieldlab.de":                                   "Yieldlab",
	"yieldlab.net":                                  "Yieldlab",
	"ad3.io":                                        "Ad3media",
	"ad6media.es":                                   "Ad6Media",
	"ad6media.fr":                                   "Ad6Media",
	"www.ad6media.fr":                               "Ad6Media",
	"adbistro.com":                                  "Adbistro",
	"adcolony.com":                                  "AdColony",
	"adingo.jp":                                     "Fluct",
	"adingo.jp<http://adingo.jp>":                   "Fluct",
	"admanmedia.com":                                "Adman Media",
	"admedia.com":                                   "AdMedia",
	"admixer.com":                                   "AdMixer",
	"admixer.net":                                   "AdMixer",
	"ads.stickyadstv.com":                           "FreeWheel",
	"ads4pics.com":                                  "Ads4Pics",
	"adtech.com<http://adtech.com>":                 "One by AOL: Display",
	"aolcloud.com":                                  "One by AOL: Display",
	"aolcloud.net<http://aolcloud.net>":             "One by AOL: Display",
	"adunity.com":                                   "Adunity",
	"advbo.ammadv.it":                               "AMM Media Marketing",
	"Advertise.com":                                 "Advertise.com",
	"advertising.com<http://advertising.com>":       "One by AOL: Video",
	"aerserv.com":                                   "Aerserv",
	"andbeyond.media":                               "AndBeyond.Media",
	"app.tv":                                        "appTV",
	"apptv.com":                                     "appTV",
	"aralego.com":                                   "ucfunnel",
	"atemda.com":                                    "WideOrbit",
	"aximusag":                                      "Aximus",
	"aximus.ch":                                     "Aximus",
	"baronsmedia.com":                               "BaronsMedia",
	"bidsxchange.com":                               "Streamlyn",
	"bidtheatre.com":                                "Bidtheater",
	"buysellads.com":                                "Buy Sell Ads",
	"carambo.la":                                    "Carambola",
	"carambola.com":                                 "Carambola",
	"cedato.com":                                    "Cedato",
	"clickio.com":                                   "Clickio",
	"collectiveuk.com":                              "Collective",
	"connectignite.com":                             "Adimia",
	"converge-digital.com":                          "Converge-Digital",
	"crimtan.com":                                   "Crimtan",
	"defymedia.com":                                 "Defy",
	"distrcitm.io":                                  "District M",
	"districtmadexchange":                           "District M",
	"districtm":                                     "District M",
	"districtm.net":                                 "District M",
	"districtmio.com":                               "District M",
	"distroscale.com":                               "DistroScale",
	"dynadmic":                                      "DynAdmic",
	"e-planning.net":                                "e-Planning",
	"eadv.it":                                       "EADV",
	"easyplatform.com":                              "Easy Platform",
	"eboundservices.com":                            "eBoundServices",
	"electric-sheep.tv":                             "Electric Sheep",
	"firstimpression.io":                            "FirstImpression.io",
	"geekexchange.com":                              "Exclude",
	"getintent.com":                                 "Get Intent",
	"glucompany.com":                                "Glu Company",
	"gmossp.jp":                                     "GMO SSP",
	"gobrowsi.com":                                  "Browsi",
	"gourmetads.com":                                "Gourmet Ads",
	"hiro-media.com":                                "Hiro Media",
	"ibillboard.com":                                "iBillboard",
	"increaserev.com":                               "Increase Rev",
	"infolinks.com":                                 "Infolinks",
	"insticator.com":                                "Insticator",
	"justpremium.com":                               "JustPremium",
	"jwdemandadexchange":                            "JWPlayer",
	"keenkale.com":                                  "KeenKale",
	"lifestreet.com":                                "Lifestreet",
	"linicom":                                       "Linicom",
	"madadsmedia.com":                               "MadAdsMedia",
	"mediabong.net":                                 "Vuble",
	"mediadeguate.com":                              "Deguate",
	"memevideoad.com":                               "MemeGlobal",
	"stinger.memeglobal.com":                        "MemeGlobal",
	"mgid.com":                                      "Mgid",
	"monarchads.com":                                "Monarch Ads",
	"netseer.com":                                   "Netseer",
	"oogle.com":                                     "Google",
	"ooyala.com":                                    "Ooyala",
	"optimatic.com":                                 "Optimatic",
	"padsquad.com":                                  "Padsquad",
	"paypal.com":                                    "Paypal",
	"playtouch":                                     "Playtouch",
	"playtouch2":                                    "Playtouch",
	"playwire.com":                                  "Paywire",
	"powerlinks.com":                                "PowerLinks",
	"pubgears.com":                                  "NexTag",
	"purch.com":                                     "Purch",
	"q1media.com":                                   "Q1 Media",
	"quantcast.com":                                 "Quantcast",
	"quantum-advertising.com":                       "Quantum Native",
	"reklamstore.com":                               "ReklamStore",
	"rekmob.com":                                    "RekMob",
	"smartadserver:smartadserver.com":               "Smart RTB+",
	"smartadsever.com":                              "Smart RTB+",
	"smartclip.net":                                 "Smartclip",
	"smartyads.com":                                 "Smarty Ads",
	"somoaudience.com":                              "Somo Audience",
	"SpotIM":                                        "Spot.im",
	"sprout-ad.com":                                 "Sprout",
	"ssphwy.com":                                    "SSPHwy",
	"startapp.com":                                  "StartApp",
	"synapsys.us":                                   "SNT Media",
	"tabletmedia.co.uk":                             "TabletMedia",
	"tappx.com":                                     "Tappx",
	"themoneytizer.com":                             "The Moneytizer",
	"thetradedesk.com":                              "The Trade Desk",
	"thrive.plus":                                   "Thrive",
	"tisoomi-services.com":                          "Tisoomi",
	"tribalfusion.com":                              "Tribal Fusion",
	"trion.com":                                     "Trion Interactive",
	"trioninteractive.com":                          "Trion Interactive",
	"truex.com":                                     "TrueX",
	"turf.digital":                                  "Turf Digital",
	"ubm.com":                                       "UBM",
	"udmserve.net":                                  "Underdog Media",
	"valueclickmedia.com":                           "Alliance Data",
	"vertamedia.com":                                "Verta Media",
	"vertoz.com":                                    "Vertoz",
	"vi.ai":                                         "Video Intelligence",
	"www.vi.ai":                                     "Video Intelligence",
	"x.fidelity-media.com":                          "Fidelity Media",
	"yandex.ru":                                     "Yandex",
	"yellowhammer.com":                              "Yellow Hammer",
	"rockyou.com":                                   "RockYou",
	"rockyou.net":                                   "RockYou",
	"innity.com":                                    "Innity",
	"innity.net":                                    "Innity",
	"advenueplatform.com":                           "Innity",
	"nativeads.com":                                 "Native Ads",
	"natiiveads.com":                                "Native Ads",
	"richaudience.com":                              "RichAudience",
	"adstanding.com":                                "AdStanding",
	"www.mass2.com":                                 "Mass2",
	"RTK.io":                                        "RTK.io",
	"atomx.com":                                     "Atomx",
	"ato.mx":                                        "Atomx",
	"rtb.ato.mx":                                    "Atomx",
	"p.ato.mx":                                      "Atomx",
	"addroplet.com":                                 "Addroplet.com ",
	"Liondigitalserving.com":                        "Liondigitalserving.com",
	"sulvo.com":                                     "sulvo.com",
	"surgeprice.com":                                "surgeprice.com",
	"mediabong.com":                                 "mediabong.com",
	"babaroll.com":                                  "Seracast",
	"juicenectar.com":                               "Juice Nectar",
}

var exnames_canonical = map[string]string{
	"Addroplet.com ":                  "addroplet.com ",
	"Google AdX":                      "google.com",
	"Adform":                          "adform.net",
	"Axonix":                          "axonix.net",
	"BidSwitch":                       "bidswitch.com",
	"BrightRoll Exchange for Display": "btrll.com",
	"C1 Exchange":                     "c1exchange.com",
	"Cox Media Technology":            "coxmt.com",
	"E-Planning":                      "e-planning.net",
	"Geniee":                          "geniee.co.jp",
	"MediaMath Curated Market":        "mm-curated",
	"Microsoft Advertising Exchange":  "aol.com",
	"MoPub via BidSwitch":             "bidswitch.com",
	"ONE by AOL: Display MP":          "adtech.com",
	"ONE by AOL: Mobile":              "aol.com",
	"OneTag":                          "onetag",
	"Smart RTB":                       "smartadserver.com",
	"Stroer SSP":                      "stroeer.com",
	"YIELD ONE":                       "platform-one.co.jp",
	"33Across":                        "33across.com",
	"ADMIZED":                         "admized.com",
	"AMM Media Marketing":             "advbo.ammadv.it",
	"Ad3media":                        "ad3media.com",
	"Ad6Media":                        "ad6media.fr",
	"AdColony":                        "adcolony.com",
	"AdForm":                          "adform.com",
	"AdMedia":                         "admedia.com",
	"AdMixer":                         "admixer.com",
	"AdStanding":                      "adstanding.com",
	"AdYouLike":                       "adyoulike.com",
	"Adbistro":                        "adbistro.com",
	"Addroplet.com":                   "addroplet.com",
	"Adfrontiers":                     "adfrontiers.com",
	"Adimia":                          "connectignite.com",
	"Adman Media":                     "admanmedia.com",
	"Ads4Pics":                        "ads4pics.com",
	"Adstir":                          "ad-stir.com",
	"Adunity":                         "adunity.com",
	"Advertise.com":                   "advertise.com",
	"Aerserv":                         "aerserv.com",
	"Algovid":                         "c.algovid.com",
	"Alliance Data":                   "valueclickmedia.com",
	"Altitude Digital":                "altitudedigital.com",
	"Amazon":                          "amazon-adsystem.com",
	"AndBeyond.Media":                 "andbeyond.media",
	"AppNexus":                        "appnexus.com",
	"Atomx":                           "atomx.com",
	"Aximus":                          "aximus.ch",
	"BaronsMedia":                     "baronsmedia.com",
	"Bidfluence":                      "bidfluence.com",
	"Bidtellect":                      "bidtellect.com",
	"Bidtheater":                      "bidtheatre.com",
	"BrightRoll from Yahoo!":          "btrll.com",
	"Brightcom":                       "brightcom.com",
	"Browsi":                          "gobrowsi.com",
	"Buy Sell Ads":                    "buysellads.com",
	"C1X":                             "c1exchange.com",
	"COMET":                           "coxmt.com",
	"Carambola":                       "carambola.com",
	"Cedato":                          "cedato.com",
	"Centro Brand Exchange":           "t.brand-server.com",
	"Clickio":                         "clickio.com",
	"Collective":                      "collectiveuk.com",
	"Connatix":                        "connatix.com",
	"Converge-Digital":                "converge-digital.com",
	"Conversant":                      "conversantmedia.com",
	"Crimtan":                         "crimtan.com",
	"Criteo":                          "criteo.net",
	"Defy":                            "defymedia.com",
	"Deguate":                         "mediadeguate.com",
	"District M":                      "districtm.com",
	"DistroScale":                     "distroscale.com",
	"DynAdmic":                        "dynadmic.com",
	"EADV":                            "eadv.it",
	"Easy Platform":                   "easyplatform.com",
	"Electric Sheep":                  "electric-sheep.tv",
	"Exclude":                         "geekexchange.com",
	"Exponential":                     "exponential.com",
	"Facebook":                        "facebook.com",
	"Fidelity Media":                  "x.fidelity-media.com",
	"FirstImpression.io":              "firstimpression.io",
	"Fluct":                           "adingo.jp",
	"FreeWheel":                       "freewheel.tv",
	"Fyber":                           "fyber.com",
	"GMO SSP":                         "gmossp.jp",
	"Gemini from Yahoo!":              "yahoo.com",
	"Get Intent":                      "getintent.com",
	"Glu Company":                     "glucompany.com",
	"Google":                          "google.com",
	"Gourmet Ads":                     "gourmetads.com",
	"GumGum":                          "gumgum.com",
	"Hiro Media":                      "hiro-media.com",
	"Imonomy":                         "imonomy.com",
	"Improve Digital":                 "improvedigital.com",
	"Increase Rev":                    "increaserev.com",
	"Index Exchange":                  "indexexchange.com",
	"Infolinks":                       "infolinks.com",
	"Inneractive":                     "inner-active.com",
	"Innity":                          "innity.com",
	"Insticator":                      "insticator.com",
	"JWPlayer":                        "jwdemandadexchange",
	"Juice Nectar":                    "juicenectar.com",
	"JustPremium":                     "justpremium.com",
	"Kargo":                           "kargo.com",
	"KeenKale":                        "keenkale.com",
	"Kiosked":                         "kiosked.com",
	"Kixer":                           "kixer.com",
	"Komoona":                         "komoonaltd.com",
	"LKQD":                            "lkqd.com",
	"Lifestreet":                      "lifestreet.com",
	"Linicom":                         "lini.com",
	"Liondigitalserving.com":          "liondigitalserving.com",
	"LiveIntent":                      "liveintent.com",
	"Lockerdome":                      "lockerdome.com",
	"LoopMe":                          "loopme.com",
	"MadAdsMedia":                     "madadsmedia.com",
	"Mass2":                           "www.mass2.com",
	"Media.net":                       "media.net",
	"MediaMath":                       "mediamath.com",
	"MemeGlobal":                      "memeglobal.com",
	"Mgid":                            "mgid.com",
	"MoPub":                           "mopub.com",
	"Mobfox":                          "mobfox.com",
	"Monarch Ads":                     "monarchads.com",
	"Native Ads":                      "nativeads.com",
	"Nativo":                          "nativo.com",
	"Netseer":                         "netseer.com",
	"NexTag":                          "pubgears.com",
	"OMS":                             "oms.eu",
	"One by AOL: Display": "adtech.com",
	"One by AOL: Mobile":  "aol.com",
	"One by AOL: Video":   "advertising.com",
	"Ooyala":              "ooyala.com",
	"OpenX":               "openx.com",
	"Optimatic":           "optimatic.com",
	"Outbrain":            "outbrain.com",
	"Padsquad":            "padsquad.com",
	"Paypal":              "paypal.com",
	"Paywire":             "playwire.com",
	"PixFuture":           "pixfuture.com",
	"Playtouch":           "playtouch.com",
	"Positive Mobile":     "positivemobile.com",
	"PowerInbox":          "rs-stripe.com",
	"PowerLinks":          "powerlinks.com",
	"PubMatic":            "pubmatic.com",
	"PulsePoint":          "contextweb.com",
	"Purch":               "purch.com",
	"Q1 Media":            "q1media.com",
	"Quantcast":           "quantcast.com",
	"Quantum Native":      "quantum-advertising.com",
	"RTK.io":              "rtk.io",
	"RekMob":              "rekmob.com",
	"ReklamStore":         "reklamstore.com",
	"RevContent":          "revcontent.com",
	"RhythmOne":           "rhythmone.com",
	"RichAudience":        "richaudience.com",
	"RockYou":             "rockyou.net",
	"Rubicon Project":     "rubiconproject.com",
	"SNT Media":           "synapsys.us",
	"SSPHwy":              "ssphwy.com",
	"Sekindo":             "sekindo.com",
	"Seracast":            "babaroll.com",
	"Sharethrough":        "sharethrough.com",
	"SlimCut Media":       "freeskreen.com",
	"Smaato":              "smaato.com",
	"Smart RTB+":          "smartadserver.com",
	"SmartStream":         "smartstream.tv",
	"Smartclip":           "smartclip.net",
	"Smarty Ads":          "smartyads.com",
	"Somo Audience":       "somoaudience.com",
	"Sonobi":              "sonobi.com",
	"Sortable":            "deployads.com",
	"Sovrn":               "sovrn.com",
	"SpotX":               "spotx.tv",
	"Spoutable":           "spoutable.com",
	"SpringServe":         "springserve.com",
	"Sprout":              "sprout-ad.com",
	"StartApp":            "startapp.com",
	"StreamRail":          "streamrail.net",
	"Streamlyn":           "bidsxchange.com",
	"Stroer":              "stroeer.com",
	"Switch Concepts":     "switchconcepts.com",
	"Synacor":             "synacor.com",
	"TabletMedia":         "tabletmedia.co.uk",
	"Taboola":             "taboola.com",
	"Tappx":               "tappx.com",
	"Teads":               "teads.tv",
	"Technorati":          "technorati.com",
	"Telaria":             "tremorhub.com",
	"The Moneytizer":      "themoneytizer.com",
	"The Trade Desk":      "thetradedesk.com",
	"Thrive":              "thrive.plus",
	"TidalTV":             "tidaltv.com",
	"Tisoomi":             "tisoomi-services.com",
	"Tribal Fusion":       "tribalfusion.com",
	"Trion Interactive":   "trioninteractive.com",
	"TripleLift":          "triplelift.com",
	"TrueX":               "truex.com",
	"TrustX":              "trustx.org",
	"Turf Digital":        "turf.digital",
	"Twiago":              "twiago.com",
	"UBM":                 "ubm.com",
	"Underdog Media":      "udmserve.net",
	"Undertone":           "undertone.com",
	"UnrulyX":             "video.unrulymedia.com",
	"Verta Media":         "vertamedia.com",
	"Vertoz":              "vertoz.com",
	"Vidazoo":             "vidazoo.com",
	"Video Intelligence":  "vi.ai",
	"Videoflare":          "videoflare.com",
	"Videology":           "videologygroup.com",
	"Vuble":               "mediabong.com",
	"WideOrbit":           "atemda.com",
	"Widespace":           "widespace.com",
	"Xapads":              "xapads.com",
	"Yandex":              "yandex.ru",
	"Yellow Hammer":       "yellowhammer.com",
	"Yieldbot":            "yldbt.com",
	"Yieldlab":            "yieldlab.net",
	"Yieldmo":             "yieldmo.com",
	"YuMe":                "yume.com",
	"Zedo":                "zedo.com",
	"appTV":               "app.tv",
	"bRealtime":           "brealtime.com",
	"e-Planning":          "e-planning.net",
	"eBoundServices":      "eboundservices.com",
	"iBillboard":          "ibillboard.com",
	"mediabong.com":       "mediabong.com",
	"sulvo.com":           "sulvo.com",
	"surgeprice.com":      "surgeprice.com",
	"ucfunnel":            "aralego.com",
}
