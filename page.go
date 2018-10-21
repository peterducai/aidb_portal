package main

//HTMLTable table in html with data
type HTMLTable struct {
	header string
	rows   string
	footer string
}

//HEADER of webpage
const HEADER = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="utf-8">

  <title>The HTML5 Herald</title>
  <meta name="description" content="The HTML5 Herald">
  <meta name="author" content="SitePoint">

  <link rel="stylesheet" href="css/styles.css?v=1.0">

  <!--[if lt IE 9]>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/html5shiv/3.7.3/html5shiv.js"></script>
  <![endif]-->
</head>

<body>
`

//FOOTER of webpage
const FOOTER = `
</body>
</html>
`

// type Page struct {
// 	//MiligramScript  = `*,*:after,*:before{box-sizing:inherit}html{box-sizing:border-box;font-size:62.5%}body{color:#606c76;font-family:'Roboto', 'Helvetica Neue', 'Helvetica', 'Arial', sans-serif;font-size:1.6em;font-weight:300;letter-spacing:.01em;line-height:1.6}blockquote{border-left:0.3rem solid #d1d1d1;margin-left:0;margin-right:0;padding:1rem 1.5rem}blockquote *:last-child{margin-bottom:0}.button,button,input[type='button'],input[type='reset'],input[type='submit']{background-color:#9b4dca;border:0.1rem solid #9b4dca;border-radius:.4rem;color:#fff;cursor:pointer;display:inline-block;font-size:1.1rem;font-weight:700;height:3.8rem;letter-spacing:.1rem;line-height:3.8rem;padding:0 3.0rem;text-align:center;text-decoration:none;text-transform:uppercase;white-space:nowrap}.button:focus,.button:hover,button:focus,button:hover,input[type='button']:focus,input[type='button']:hover,input[type='reset']:focus,input[type='reset']:hover,input[type='submit']:focus,input[type='submit']:hover{background-color:#606c76;border-color:#606c76;color:#fff;outline:0}.button[disabled],button[disabled],input[type='button'][disabled],input[type='reset'][disabled],input[type='submit'][disabled]{cursor:default;opacity:.5}.button[disabled]:focus,.button[disabled]:hover,button[disabled]:focus,button[disabled]:hover,input[type='button'][disabled]:focus,input[type='button'][disabled]:hover,input[type='reset'][disabled]:focus,input[type='reset'][disabled]:hover,input[type='submit'][disabled]:focus,input[type='submit'][disabled]:hover{background-color:#9b4dca;border-color:#9b4dca}.button.button-outline,button.button-outline,input[type='button'].button-outline,input[type='reset'].button-outline,input[type='submit'].button-outline{background-color:transparent;color:#9b4dca}.button.button-outline:focus,.button.button-outline:hover,button.button-outline:focus,button.button-outline:hover,input[type='button'].button-outline:focus,input[type='button'].button-outline:hover,input[type='reset'].button-outline:focus,input[type='reset'].button-outline:hover,input[type='submit'].button-outline:focus,input[type='submit'].button-outline:hover{background-color:transparent;border-color:#606c76;color:#606c76}.button.button-outline[disabled]:focus,.button.button-outline[disabled]:hover,button.button-outline[disabled]:focus,button.button-outline[disabled]:hover,input[type='button'].button-outline[disabled]:focus,input[type='button'].button-outline[disabled]:hover,input[type='reset'].button-outline[disabled]:focus,input[type='reset'].button-outline[disabled]:hover,input[type='submit'].button-outline[disabled]:focus,input[type='submit'].button-outline[disabled]:hover{border-color:inherit;color:#9b4dca}.button.button-clear,button.button-clear,input[type='button'].button-clear,input[type='reset'].button-clear,input[type='submit'].button-clear{background-color:transparent;border-color:transparent;color:#9b4dca}.button.button-clear:focus,.button.button-clear:hover,button.button-clear:focus,button.button-clear:hover,input[type='button'].button-clear:focus,input[type='button'].button-clear:hover,input[type='reset'].button-clear:focus,input[type='reset'].button-clear:hover,input[type='submit'].button-clear:focus,input[type='submit'].button-clear:hover{background-color:transparent;border-color:transparent;color:#606c76}.button.button-clear[disabled]:focus,.button.button-clear[disabled]:hover,button.button-clear[disabled]:focus,button.button-clear[disabled]:hover,input[type='button'].button-clear[disabled]:focus,input[type='button'].button-clear[disabled]:hover,input[type='reset'].button-clear[disabled]:focus,input[type='reset'].button-clear[disabled]:hover,input[type='submit'].button-clear[disabled]:focus,input[type='submit'].button-clear[disabled]:hover{color:#9b4dca}code{background:#f4f5f6;border-radius:.4rem;font-size:86%;margin:0 .2rem;padding:.2rem .5rem;white-space:nowrap}pre{background:#f4f5f6;border-left:0.3rem solid #9b4dca;overflow-y:hidden}pre>code{border-radius:0;display:block;padding:1rem 1.5rem;white-space:pre}hr{border:0;border-top:0.1rem solid #f4f5f6;margin:3.0rem 0}input[type='email'],input[type='number'],input[type='password'],input[type='search'],input[type='tel'],input[type='text'],input[type='url'],textarea,select{-webkit-appearance:none;-moz-appearance:none;appearance:none;background-color:transparent;border:0.1rem solid #d1d1d1;border-radius:.4rem;box-shadow:none;box-sizing:inherit;height:3.8rem;padding:.6rem 1.0rem;width:100%}input[type='email']:focus,input[type='number']:focus,input[type='password']:focus,input[type='search']:focus,input[type='tel']:focus,input[type='text']:focus,input[type='url']:focus,textarea:focus,select:focus{border-color:#9b4dca;outline:0}select{background:url('data:image/svg+xml;utf8,<svg xmlns="http://www.w3.org/2000/svg" height="14" viewBox="0 0 29 14" width="29"><path fill="#d1d1d1" d="M9.37727 3.625l5.08154 6.93523L19.54036 3.625"/></svg>') center right no-repeat;padding-right:3.0rem}select:focus{background-image:url('data:image/svg+xml;utf8,<svg xmlns="http://www.w3.org/2000/svg" height="14" viewBox="0 0 29 14" width="29"><path fill="#9b4dca" d="M9.37727 3.625l5.08154 6.93523L19.54036 3.625"/></svg>')}textarea{min-height:6.5rem}label,legend{display:block;font-size:1.6rem;font-weight:700;margin-bottom:.5rem}fieldset{border-width:0;padding:0}input[type='checkbox'],input[type='radio']{display:inline}.label-inline{display:inline-block;font-weight:normal;margin-left:.5rem}.container{margin:0 auto;max-width:112.0rem;padding:0 2.0rem;position:relative;width:100%}.row{display:flex;flex-direction:column;padding:0;width:100%}.row.row-no-padding{padding:0}.row.row-no-padding>.column{padding:0}.row.row-wrap{flex-wrap:wrap}.row.row-top{align-items:flex-start}.row.row-bottom{align-items:flex-end}.row.row-center{align-items:center}.row.row-stretch{align-items:stretch}.row.row-baseline{align-items:baseline}.row .column{display:block;flex:1 1 auto;margin-left:0;max-width:100%;width:100%}.row .column.column-offset-10{margin-left:10%}.row .column.column-offset-20{margin-left:20%}.row .column.column-offset-25{margin-left:25%}.row .column.column-offset-33,.row .column.column-offset-34{margin-left:33.3333%}.row .column.column-offset-50{margin-left:50%}.row .column.column-offset-66,.row .column.column-offset-67{margin-left:66.6666%}.row .column.column-offset-75{margin-left:75%}.row .column.column-offset-80{margin-left:80%}.row .column.column-offset-90{margin-left:90%}.row .column.column-10{flex:0 0 10%;max-width:10%}.row .column.column-20{flex:0 0 20%;max-width:20%}.row .column.column-25{flex:0 0 25%;max-width:25%}.row .column.column-33,.row .column.column-34{flex:0 0 33.3333%;max-width:33.3333%}.row .column.column-40{flex:0 0 40%;max-width:40%}.row .column.column-50{flex:0 0 50%;max-width:50%}.row .column.column-60{flex:0 0 60%;max-width:60%}.row .column.column-66,.row .column.column-67{flex:0 0 66.6666%;max-width:66.6666%}.row .column.column-75{flex:0 0 75%;max-width:75%}.row .column.column-80{flex:0 0 80%;max-width:80%}.row .column.column-90{flex:0 0 90%;max-width:90%}.row .column .column-top{align-self:flex-start}.row .column .column-bottom{align-self:flex-end}.row .column .column-center{-ms-grid-row-align:center;align-self:center}@media (min-width: 40rem){.row{flex-direction:row;margin-left:-1.0rem;width:calc(100% + 2.0rem)}.row .column{margin-bottom:inherit;padding:0 1.0rem}}a{color:#9b4dca;text-decoration:none}a:focus,a:hover{color:#606c76}dl,ol,ul{list-style:none;margin-top:0;padding-left:0}dl dl,dl ol,dl ul,ol dl,ol ol,ol ul,ul dl,ul ol,ul ul{font-size:90%;margin:1.5rem 0 1.5rem 3.0rem}ol{list-style:decimal inside}ul{list-style:circle inside}.button,button,dd,dt,li{margin-bottom:1.0rem}fieldset,input,select,textarea{margin-bottom:1.5rem}blockquote,dl,figure,form,ol,p,pre,table,ul{margin-bottom:2.5rem}table{border-spacing:0;width:100%}td,th{border-bottom:0.1rem solid #e1e1e1;padding:1.2rem 1.5rem;text-align:left}td:first-child,th:first-child{padding-left:0}td:last-child,th:last-child{padding-right:0}b,strong{font-weight:bold}p{margin-top:0}h1,h2,h3,h4,h5,h6{font-weight:300;letter-spacing:-.1rem;margin-bottom:2.0rem;margin-top:0}h1{font-size:4.6rem;line-height:1.2}h2{font-size:3.6rem;line-height:1.25}h3{font-size:2.8rem;line-height:1.3}h4{font-size:2.2rem;letter-spacing:-.08rem;line-height:1.35}h5{font-size:1.8rem;letter-spacing:-.05rem;line-height:1.5}h6{font-size:1.6rem;letter-spacing:0;line-height:1.4}img{max-width:100%}.clearfix:after{clear:both;content:' ';display:table}.float-left{float:left}.float-right{float:right}`
// }

// //Header of page
// const Header = `<!doctype html>
// <html lang="en">
// 	<head>
// 		<meta charset="utf-8">
// 		<meta name="viewport" content="width=device-width, initial-scale=1">
// 		<meta name="author" content="peter daemonna ducai">
// 		<meta name="description" content="your admin for asset inventory.">
// 		<title>AIDB Portal | your admin for asset inventory.</title>
// 		<link rel="icon" href="https://milligram.github.io/images/icon.png">
// 		<link rel="stylesheet" href="roboto.woff2">
// 		<link rel="stylesheet" href="normalize.css">
// 		<link rel="stylesheet" href="milligram.min.css">
// 		<link rel="stylesheet" href="main.css">
// 	</head>
// 	<body>`

// //Footer of page
// const Footer = `			<footer class="footer">
// 	<section class="container">
// 		<p>Designed with ♥ by <a target="_blank" href="https://github.com/peterducai/aidb_portal" title="AIDB Portal">AIDB</a>. Licensed under the <a target="_blank" href="https://github.com/milligram/milligram#license" title="Apache2 License">Apache2 License</a>.</p>
// 	</section>
// </footer>

// </main>

// <script src="https://milligram.github.io/scripts/main.js"></script>

// </body>
// </html>`

// //Content of page
// const Content = `<main class="wrapper">
// <nav class="navigation">
// 	<section class="container">
// 		<a class="navigation-title" href="https://github.com/peterducai/aidb_portal">
// 			<svg class="img" version="1.1" viewBox="0 0 463 669">
// 				<g transform="translate(0.000000,669.000000) scale(0.100000,-0.100000)">
// 					<path d="M2303 6677c-11-13-58-89-393-627-128-206-247-397-265-425-18-27-85-135-150-240-65-104-281-451-480-770-358-575-604-970-641-1032-10-18-45-74-76-126-47-78-106-194-107-212-1-3-11-26-24-53-60-118-132-406-157-623-19-158-8-491 20-649 82-462 291-872 619-1213 192-199 387-340 646-467 335-165 638-235 1020-235 382 0 685 70 1020 235 259 127 454 268 646 467 328 341 537 751 619 1213 28 158 39 491 20 649-25 217-97 505-157 623-13 27-23 50-23 53 0 16-57 127-107 210-32 52-67 110-77 128-37 62-283 457-641 1032-199 319-415 666-480 770-65 105-132 213-150 240-18 28-137 219-265 425-354 570-393 630-400 635-4 3-12-1-17-8zm138-904c118-191 654-1050 1214-1948 148-236 271-440 273-452 2-13 8-23 11-23 14 0 72-99 125-212 92-195 146-384 171-598 116-974-526-1884-1488-2110-868-205-1779 234-2173 1046-253 522-257 1124-10 1659 45 97 108 210 126 225 4 3 9 13 13 22 3 9 126 209 273 445 734 1176 1102 1766 1213 1946 67 108 124 197 126 197 2 0 59-89 126-197zM1080 3228c-75-17-114-67-190-243-91-212-128-368-137-580-34-772 497-1451 1254-1605 77-15 112-18 143-11 155 35 212 213 106 329-32 36-62 48-181 75-223 50-392 140-552 291-115 109-178 192-242 316-101 197-136 355-128 580 3 111 10 167 30 241 30 113 80 237 107 267 11 12 20 26 20 32 0 6 8 22 17 36 26 41 27 99 3 147-54 105-142 149-250 125z"></path>
// 				</g>
// 			</svg>
// 			&nbsp;
// 			<h1 class="title">AIDB Portal</h1>
// 		</a>
// 		<ul class="navigation-list float-right">
// 			<li class="navigation-item">
// 				<a class="navigation-link" href="#popover-grid" data-popover>Docs</a>
// 				<div class="popover" id="popover-grid">
// 					<ul class="popover-list">
// 						<li class="popover-item"><a class="popover-link" href="#getting-started" title="Getting Started">Getting Started</a></li>
// 						<li class="popover-item"><a class="popover-link" href="#typography" title="Typography">Typography</a></li>
// 						<li class="popover-item"><a class="popover-link" href="#blockquotes" title="Blockquotes">Blockquotes</a></li>
// 						<li class="popover-item"><a class="popover-link" href="#buttons" title="Buttons">Buttons</a></li>
// 						<li class="popover-item"><a class="popover-link" href="#lists" title="Lists">Lists</a></li>
// 						<li class="popover-item"><a class="popover-link" href="#forms" title="Forms">Forms</a></li>
// 						<li class="popover-item"><a class="popover-link" href="#tables" title="Tables">Tables</a></li>
// 						<li class="popover-item"><a class="popover-link" href="#grids" title="Grids">Grids</a></li>
// 						<li class="popover-item"><a class="popover-link" href="#codes" title="Codes">Codes</a></li>
// 						<li class="popover-item"><a class="popover-link" href="#utilities" title="Utilities">Utilities</a></li>
// 						<li class="popover-item"><a class="popover-link" href="#tips" title="Tips">Tips</a></li>
// 						<li class="popover-item"><a class="popover-link" href="#browser-support" title="Browser Support">Browser Support</a></li>
// 						<li class="popover-item"><a class="popover-link" href="#examples" title="Examples">Examples</a></li>
// 						<li class="popover-item"><a class="popover-link" href="#contributing" title="Contributing">Contributing</a></li>
// 					</ul>
// 				</div>
// 			</li>
// 			<li class="navigation-item">
// 				<a class="navigation-link" href="#popover-support" data-popover>Support</a>
// 				<div class="popover" id="popover-support">
// 					<ul class="popover-list">
// 						<li class="popover-item"><a class="popover-link" target="blank" href="https://github.com/milligram/milligram" title="On Github">On Github</a></li>
// 						<li class="popover-item"><a class="popover-link" target="blank" href="https://codepen.io/milligramcss" title="On Codepen">On Codepen</a></li>
// 						<li class="popover-item"><a class="popover-link" target="blank" href="https://facebook.com/milligramcss" title="On Facebook">On Facebook</a></li>
// 						<li class="popover-item"><a class="popover-link" target="blank" href="https://twitter.com/milligramcss" title="On Twitter">On Twitter</a></li>
// 						<li class="popover-item"><a class="popover-link" target="blank" href="https://github.com/milligram/milligram/issues/new" title="Need help?">Need help?</a></li>
// 						<li class="popover-item"><a class="popover-link" target="blank" href="https://github.com/milligram/milligram#license" title="License">License</a></li>
// 						<li class="popover-item"><a class="popover-link" target="blank" href="https://github.com/milligram/milligram/releases" title="Versions">Versions</a></li>
// 					</ul>
// 				</div>
// 			</li>
// 		</ul>
// 		<a href="https://github.com/peterducai/aidb" title="AIDB on Github" target="_blank">
// 			<svg class="octocat" viewBox="0 0 250 250"><path d="M0,0 L115,115 L130,115 L142,142 L250,250 L250,0 Z"></path><path class="octocat-arm" d="M128.3,109.0 C113.8,99.7 119.0,89.6 119.0,89.6 C122.0,82.7 120.5,78.6 120.5,78.6 C119.2,72.0 123.4,76.3 123.4,76.3 C127.3,80.9 125.5,87.3 125.5,87.3 C122.9,97.6 130.6,101.9 134.4,103.2"></path><path class="octocat-body" d="M115.0,115.0 C114.9,115.1 118.7,116.5 119.8,115.4 L133.7,101.6 C136.9,99.2 139.9,98.4 142.2,98.6 C133.8,88.0 127.5,74.4 143.8,58.0 C148.5,53.4 154.0,51.2 159.7,51.0 C160.3,49.4 163.2,43.6 171.4,40.1 C171.4,40.1 176.1,42.5 178.8,56.2 C183.1,58.6 187.2,61.8 190.9,65.4 C194.5,69.0 197.7,73.2 200.1,77.6 C213.8,80.2 216.3,84.9 216.3,84.9 C212.7,93.1 206.9,96.0 205.4,96.6 C205.1,102.4 203.0,107.8 198.3,112.5 C181.9,128.9 168.3,122.5 157.7,114.1 C157.9,116.9 156.7,120.9 152.7,124.9 L141.0,136.5 C139.8,137.7 141.6,141.9 141.8,141.8 Z"></path></svg>
// 		</a>
// 	</section>
// </nav>
// <section class="container" id="examples">
// 	<h5 class="title">Welcome to main page</h5>
// 	<p>
// 	<p><b>LOCATIONS</b></p>
// 	<p>
// 	<table>
// 	<thead>
// 	  <tr>
// 		<th>ID</th>
// 		<th>Country</th>
// 		<th>Town</th>
// 		<th>Address</th>
// 		<th>Device count</th>
// 	  </tr>
// 	</thead>
// 	<tbody>
// 	  <tr>
// 		<td>US7</td>
// 		<td>USA</td>
// 		<td>Akron, OH</td>
// 		<td>Av 18453</td>
// 		<td>3527</td>
// 	  </tr>
// 	  <tr>
// 		<td>CA1</td>
// 		<td>Canada</td>
// 		<td>Quebec</td>
// 		<td>Quebec Avenue 458</td>
// 		<td>1227</td>
// 	  </tr>
// 	</tbody>
//   </table>
// 	</p>
// 	<p><b>HEALTH status</b></p>
// 	<p>
// 	<table>
// 	<thead>
// 	  <tr>
// 	    <th>Datacenter</th>
// 		<th>Active devices</th>
// 		<th>Unrecheable</th>
// 		<th>EOL</th>
// 		<th>Expired licenses</th>
// 	  </tr>
// 	</thead>
// 	<tbody>
// 	  <tr>
// 		<td>US7</td>
// 		<td>2227</td>
// 		<td>3</td>
// 		<td>25</td>
// 		<td>0</td>
// 	  </tr>
// 	  <tr>
// 		<td>CA1</td>
// 		<td>725</td>
// 		<td>3</td>
// 		<td>77</td>
// 		<td>3</td>
// 	  </tr>
// 	</tbody>
//   </table>
// 	</p>
// </section>
// <section class="container" id="contributing">
// 	<h5 class="title">Contributing</h5>
// 	<p>Want to contribute? Follow these <a href="https://github.com/milligram/milligram/blob/master/.github/contributing.md" title="Contributing">recommendations</a>.</p>
// </section>`
