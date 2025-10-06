package art

var Ship = []RawStaticArt{
	{
		Frame: `     |    |    |
    )_)  )_)  )_)
   )___))___))___)\
  )____)____)_____)\\
_____|____|____|____\\\__
\                   /`,
		Mask: `     y    y    y

                  w
                   ww
yyyyyyyyyyyyyyyyyyyywwwyy
y???????????????????y`,
	},
	{
		Frame: `         |    |    |
        (_(  (_(  (_(
      /(___((___((___(
    //(_____(____(____(
__///____|____|____|_____
    \                   /`,
		Mask: `         y    y    y

      w
    ww
yywwwyyyyyyyyyyyyyyyyyyyy
    y???????????????????y`,
	},
}

var Whale = []RawStaticArt{
	{
		Frame: `        .-----:
      .'       b.
,    /       (o) \
\b._/          ,__)`,
		Mask: `        bbbbbbb
      bb???????bb
b    b???????bwb?b
bbbbb??????????bbbb`,
	},
	{
		Frame: `    :-----.
  .'       b.
 / (o)       \    ,
(__,          \_.'/`,
		Mask: `    bbbbbbb
  bb???????bb
 b?bwb???????b    b
bbbb??????????bbbbb`,
	},
}

var WaterSpoutFrames = []string{
	`

`,
	`

   :`,
	`
   :
   :`,
	`  . .
  -:-
   :`,
	`  . .
 .-:-.
   :`,
	`  . .
'.-:-.'
'  :  '`,
	`
 .- -.
;  :  ;`,
	`

;     ;`,
}

var WaterSpoutMask = `  C C
CCCCCCC
C  C  C`

var MonsterFrames = [][]string{
	{
		`                                                          ____
            __                                          /   o  \
          /    \        _                     _       /     ____ >
  _      |  __  |     /   \        _        /   \    |     |
 | \     |  ||  |    |     |     /   \     |     |   |     |`,
		`                                                          ____
                                             __         /   o  \
             _                     _       /    \     /     ____ >
   _       /   \        _        /   \    |  __  |   |     |
  | \     |     |     /   \     |     |   |  ||  |   |     |`,
		`                                                          ____
                                  __                    /   o  \
 _                      _       /    \        _       /     ____ >
| \          _        /   \    |  __  |     /   \    |     |
 \ \       /   \     |     |   |  ||  |    |     |   |     |`,
		`                                                          ____
                       __                               /   o  \
  _          _       /    \        _                  /     ____ >
 | \       /   \    |  __  |     /   \        _      |     |
  \ \     |     |   |  ||  |    |     |     /   \    |     |`,
	},
	{
		`    ____
  /  o   \                                          __
< ____     \       _                     _        /    \
      |     |    /   \        _        /   \     |  __  |      _
      |     |   |     |     /   \     |     |    |  ||  |     / |`,
		`    ____
  /  o   \         __
< ____     \     /    \       _                     _
      |     |   |  __  |    /   \        _        /   \       _
      |     |   |  ||  |   |     |     /   \     |     |     / |`,
		`    ____
  /  o   \                    __
< ____     \       _        /    \       _                      _
      |     |    /   \     |  __  |    /   \        _          / |
      |     |   |     |    |  ||  |   |     |     /   \       / /`,
		`    ____
  /  o   \                               __
< ____     \                  _        /    \       _          _
      |     |      _        /   \     |  __  |    /   \       / |
      |     |    /   \     |     |    |  ||  |   |     |     / /`,
	},
}

var MonsterMasks = [][]string{
	{
		`                                                          gggg
            gg                                          g???w??g
          g????g        g                     g       g?????gggg?g
  g      g??gg??g     g???g        g        g???g    g?????g
 g?g     g??gg??g    g?????g     g???g     g?????g   g?????g`,
		`                                                          gggg
                                             gg         g???w??g
             g                     g       g????g     g?????gggg?g
   g       g???g        g        g???g    g??gg??g   g?????g
  g?g     g?????g     g???g     g?????g   g??gg??g   g?????g`,
		`                                                          gggg
                                  gg                    g???w??g
 g                      g       g????g        g       g?????gggg?g
g?g          g        g???g    g??gg??g     g???g    g?????g
 g?g       g???g     g?????g   g??gg??g    g?????g   g?????g`,
		`                                                          gggg
                       gg                               g???w??g
  g          g       g????g        g                  g?????gggg?g
 g?g       g???g    g??gg??g     g???g        g      g?????g
  g?g     g?????g   g??gg??g    g?????g     g???g    g?????g`,
	},
	{
		`    gggg
  g??w???g                                          gg
g?gggg?????g       g                     g        g????g
      g?????g    g???g        g        g???g     g??gg??g      g
      g?????g   g?????g     g???g     g?????g    g??gg??g     g?g`,
		`    gggg
  g??w???g         gg
g?gggg?????g     g????g       g                     g
      g?????g   g??gg??g    g???g        g        g???g       g
      g?????g   g??gg??g   g?????g     g???g     g?????g     g?g`,
		`    gggg
  g??w???g                    gg
g?gggg?????g       g        g????g       g                      g
      g?????g    g???g     g??gg??g    g???g        g          g?g
      g?????g   g?????g    g??gg??g   g?????g     g???g       g?g`,
		`    gggg
  g??w???g                               gg
g?gggg?????g                  g        g????g       g          g
      g?????g      g        g???g     g??gg??g    g???g       g?g
      g?????g    g???g     g?????g    g??gg??g   g?????g     g?g`,
	},
}

var BigFish = []RawStaticArt{
	{
		Frame: ` ______
b""-.  bbbbb-----.....__
     b.  .      .       b-.
       :     .     .       b.
 ,     :   .    .          _ :
: b.   :                  (@) b._
 b. b..'     .     =b-.       .__)
   ;     .        =  ~  :     .-"
 .' .'b.   .    .  =.-'  b._ .'
: .'   :               .   .'
 '   .'  .    .     .   .-'
   .'____....----''.'=.'
   ""             .'.'
               ''"'b`,
		Mask: ` 111111
11111??11111111111111111
     11??2??????2???????111
       1?????2?????2???????11
 1     1???2????2??????????1?1
1?11   1??????????????????1w1?111
 11?1111?????2?????1111???????1111
   1?????2????????1??1??1?????111
 11?1111???2????2??1111??111?11
1?11   1???????????????2???11
 1   11??2????2?????2???111
   111111111111111111111
   11             1111
               11111`,
	},
	{
		Frame: `                           ______
          __.....-----'''''  .-""'
       .-'       .      .  .'
     .'       .     .     :
    : _          .    .   :     ,
 _.' (@)                  :   .' :
(__.       .-'=     .     b..' .'
 "-.     :  ~  =        .     ;
   b. _.'  b-.=  .    .   .'b. b.
     b.   .               :   b. :
       b-.   .     .    .  b.   b
          b.=b.bb----....____b.
            b.b.             ""
              'b"bb`,
		Mask: `                           111111
          11111111111111111??11111
       111???????2??????2??11
     11???????2?????2?????1
    1?1??????????2????2???1     1
 111?1w1??????????????????1   11?1
1111???????1111?????2?????1111?11
 111?????1??1??1????????2?????1
   11?111??1111??2????2???1111?11
     11???2???????????????1   11?1
       111???2?????2????2??11   1
          111111111111111111111
            1111             11
              11111`,
	},
}

var Shark = []RawStaticArt{
	{
		Frame: `                              __
                             ( b\
  ,                          )   b\
;' b.                        (     b\__
 ;   b.             __..---''          b~~~~-._
  b.   b.____...--''                       (o  b--._
    >                     _.-'      .((      ._     )
  .b.-b--...__         .-'     -.___.....-(|/|/|/|/'
 ;.'         b. ...----b.___.',,,_______......---'
 '           '-'`,
		Mask: `                              cc
                             c?cc
  c                          c???cc
cc?cc                        c?????cccc
 c???cc             ccccccccc??????????cccccccc
  cc???ccccccccccccc???????????????????????cw??ccccc
    c?????????????????????cccc??????ccc??????cc?????c
  cccccccccccc?????????ccc?????ccccccccccccwwwwwwwwc
 ccc         cc?cccccccccccccccccccccccccccccccccc
 c           ccc`,
	},
	{
		Frame: `                     __
                    /' )
                  /'   (                          ,
              __/'     )                        .' b;
      _.-~~~~'          bb---..__             .'   ;
 _.--'  o)                       bb--...____.'   .'
(     _.      )).      b-._                     <
 b\|\|\|\|)-.....___.-     b-.         __...--'-.'.
   b---......_______,,,b.___.'----... .'         b.;
                                     b-b           b`,
		Mask: `                     cc
                    cc?c
                  cc???c                          c
              cccc?????c                        cc?cc
      cccccccc??????????ccccccccc             cc???c
 ccccc??wc???????????????????????ccccccccccccc???cc
c?????cc??????ccc??????cccc?????????????????????c
 cwwwwwwwwcccccccccccc?????ccc?????????cccccccccccc
   cccccccccccccccccccccccccccccccccc?cc         ccc
                                     ccc           c`,
	},
}

var SplatFrames = []string{
	`
   .
  ***
   '
`,
	`
 ",*;b
 "*,**
 *"'~'
`,
	`  , ,
 " ","'
 *" *'"
  " ; .
`,
	`* ' , ' b
' b * . '
 ' b' ",'
* ' " * .
" * ', '`,
}

var SplatMasks = []string{
	`
   r
  rrr
   r
`,
	`
 rrrrr
 rrrrr
 rrrrr
`,
	`  r r
 r rrrr
 rr rrr
  r r r
`,
	`r r r r r
r r r r r
 r rr rrr
r r r r r
r r rr r`,
}

var FishHook = RawStaticArt{
	Frame: `    O
    ||
 .   \\
/ \   ||
 \\__//
  b--'`,
	Mask: `    y
    yy
 y   yy
y?y   yy
 yyyyyy
  yyyy`,
}

// var FishHook = RawStaticArt{
// 	Frame:
// `       o
//       ||
//  .    ||
// / \   ||
//  \\__//
//   b--'`,
// 	Mask:
// `       y
//       yy
//  y    yy
// y y   yy
//  yyyyyy
//   yyyy`,
// }

var Swan = []RawStaticArt{
	{
		Frame: `       ___
,_    / _,\
| \   \( \|
|  \_  \\
(_   \_) \
(\_   b   \
 \   -=~  /`,
		Mask: `
         g
         yy



  ???   ??`,
	},
	{
		Frame: ` ___
/,_ \    _,
|/ )/   / |
  //  _/  |
 / ( /   _)
/   b   _/)
\  ~=-   /`,
		Mask: `
 g
yy



 ??   ???`,
	},
}

// Get my ducks in a row
var DuckFrames = [][]string{
	{
		`      _          _          _
,____(')=  ,____(')=  ,____(')<
 \~~= ')    \~~= ')    \~~= ')`,
		`      _          _          _
,____(')=  ,____(')<  ,____(')=
 \~~= ')    \~~= ')    \~~= ')`,
		`      _          _          _
,____(')<  ,____(')=  ,____(')=
 \~~= ')    \~~= ')    \~~= ')`,
	},
	{
		`  _          _          _
>(')____,  =(')____,  =(')____,
 (' =~~/    (' =~~/    (' =~~/`,
		`  _          _          _
=(')____,  >(')____,  =(')____,
 (' =~~/    (' =~~/    (' =~~/`,
		`  _          _          _
=(')____,  =(')____,  >(')____,
 (' =~~/    (' =~~/    (' =~~/`,
	},
}

var DuckMasks = []string{
	`      g          g          g
wwwwwgcgy??wwwwwgcgy??wwwwwgcgy
 wwww?Ww    wwww?Ww    wwww?Ww`,
	`  g          g          g
ygcgwwwww??ygcgwwwww??ygcgwwwww
 wW?wwww    wW?wwww    wW?wwww`,
}

var DolphinFrames = [][]string{
	{`        ,
      __)\_
(\_.-'    ab-.
(/~~bbbb(/~^^b`,
		`        ,
(\__  __)\_
(/~.''    ab-.
    bbbb\)~^^b`,
	},
	{`     ,
   _/(__
.-'a    b-._/)
'^^~\)''''~~\)`,
		`     ,
   _/(__  __/)
.-'a    bb.~\)
'^^~(/''''`,
	},
}

var DolphinMasks = [][][]string{
	{
		{
			`        c
      ccccc
ccccccccccwccc
cccccccccccccc`,
			`        c
cccc  ccccc
ccccccccccwccc
    cccccccccc`,
		},
		{
			`        B
      BBBBB
BBBBBBBBBBwBBB
BBBBBBBBBBBBBB`,
			`        B
BBBB  BBBBB
BBBBBBBBBBwBBB
    BBBBBBBBBB`,
		},
		{
			`        b
      bbbbb
bbbbbbbbbbwbbb
bbbbbbbbbbbbbb`,
			`        b
bbbb  bbbbb
bbbbbbbbbbwbbb
    bbbbbbbbbb`,
		},
	},
	{
		{
			`     c
   ccccc
cccwcccccccccc
cccccccccccccc`,
			`     c
   ccccc  cccc
cccwcccccccccc
cccccccccc`,
		}, {
			`     B
   BBBBB
BBBwBBBBBBBBBB
BBBBBBBBBBBBBB`,
			`     B
   BBBBB  BBBB
BBBwBBBBBBBBBB
BBBBBBBBBB`,
		}, {
			`     b
   bbbbb
bbbwbbbbbbbbbb
bbbbbbbbbbbbbb`,
			`     b
   bbbbb  bbbb
bbbwbbbbbbbbbb
bbbbbbbbbb`,
		},
	},
}
