import{_ as r}from"./plugin-vue_export-helper-c27b6911.js";import{r as i,o as d,c,d as e,e as t,b as o,w as a,f as s}from"./app-45f7c304.js";const u={},h=e("h1",{id:"jsonl-data-type-reference",tabindex:"-1"},[e("a",{class:"header-anchor",href:"#jsonl-data-type-reference","aria-hidden":"true"},"#"),t(),e("code",null,"jsonl"),t(" - Data-Type Reference")],-1),p=e("blockquote",null,[e("p",null,"JSON Lines")],-1),m=e("h2",{id:"description",tabindex:"-1"},[e("a",{class:"header-anchor",href:"#description","aria-hidden":"true"},"#"),t(" Description")],-1),_={href:"http://jsonlines.org/",target:"_blank",rel:"noopener noreferrer"},f=e("blockquote",null,[e("p",null,"JSON Lines is a convenient format for storing structured data that may be processed one record at a time. It works well with unix-style text processing tools and shell pipelines. It's a great format for log files. It's also a flexible format for passing messages between cooperating processes.")],-1),g=e("h2",{id:"examples",tabindex:"-1"},[e("a",{class:"header-anchor",href:"#examples","aria-hidden":"true"},"#"),t(" Examples")],-1),q={href:"http://jsonlines.org/examples/",target:"_blank",rel:"noopener noreferrer"},x=s(`<h3 id="tabulated-data" tabindex="-1"><a class="header-anchor" href="#tabulated-data" aria-hidden="true">#</a> Tabulated data</h3><pre><code>[&quot;Name&quot;, &quot;Session&quot;, &quot;Score&quot;, &quot;Completed&quot;]
[&quot;Gilbert&quot;, &quot;2013&quot;, 24, true]
[&quot;Alexa&quot;, &quot;2013&quot;, 29, true]
[&quot;May&quot;, &quot;2012B&quot;, 14, false]
[&quot;Deloise&quot;, &quot;2012A&quot;, 19, true]
</code></pre><p>This format is equatable to <code>generic</code> and <code>csv</code>.</p><h3 id="nested-objects" tabindex="-1"><a class="header-anchor" href="#nested-objects" aria-hidden="true">#</a> Nested objects</h3><pre><code>{&quot;name&quot;: &quot;Gilbert&quot;, &quot;wins&quot;: [[&quot;straight&quot;, &quot;7♣&quot;], [&quot;one pair&quot;, &quot;10♥&quot;]]}
{&quot;name&quot;: &quot;Alexa&quot;, &quot;wins&quot;: [[&quot;two pair&quot;, &quot;4♠&quot;], [&quot;two pair&quot;, &quot;9♠&quot;]]}
{&quot;name&quot;: &quot;May&quot;, &quot;wins&quot;: []}
{&quot;name&quot;: &quot;Deloise&quot;, &quot;wins&quot;: [[&quot;three of a kind&quot;, &quot;5♣&quot;]]}
</code></pre><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2><h3 id="concatenated-json" tabindex="-1"><a class="header-anchor" href="#concatenated-json" aria-hidden="true">#</a> Concatenated JSON</h3>`,7),y=e("code",null,"jsonl",-1),j={href:"https://en.wikipedia.org/wiki/JSON_streaming#Concatenated_JSON",target:"_blank",rel:"noopener noreferrer"},M=s(`<blockquote><p>Concatenated JSON streaming allows the sender to simply write each JSON object into the stream with no delimiters. It relies on the receiver using a parser that can recognize and emit each JSON object as the terminating character is parsed. Concatenated JSON isn&#39;t a new format, it&#39;s simply a name for streaming multiple JSON objects without any delimiters.</p><p>The advantage of this format is that it can handle JSON objects that have been formatted with embedded newline characters, e.g., pretty-printed for human readability. For example, these two inputs are both valid and produce the same output:</p><h4 id="single-line-concatenated-json" tabindex="-1"><a class="header-anchor" href="#single-line-concatenated-json" aria-hidden="true">#</a> Single line concatenated JSON</h4><pre><code>{&quot;some&quot;:&quot;thing\\n&quot;}{&quot;may&quot;:{&quot;include&quot;:&quot;nested&quot;,&quot;objects&quot;:[&quot;and&quot;,&quot;arrays&quot;]}}
</code></pre><h4 id="multi-line-concatenated-json" tabindex="-1"><a class="header-anchor" href="#multi-line-concatenated-json" aria-hidden="true">#</a> Multi-line concatenated JSON</h4><pre><code>{
  &quot;some&quot;: &quot;thing\\n&quot;
}
{
  &quot;may&quot;: {
    &quot;include&quot;: &quot;nested&quot;,
    &quot;objects&quot;: [
      &quot;and&quot;,
      &quot;arrays&quot;
    ]
  }
}
</code></pre></blockquote><p>...however in Murex&#39;s case, only single line concatenated JSON files (example 1) are supported; and that is only supported to cover some edge cases when writing JSON lines and a new line character isn&#39;t included. The primary example might be when generating JSON lines from inside a <code>for</code> loop.</p><p>This is resolved in the new data-type parser <code>jsonc</code> (Concatenated JSON). See line below.</p><h3 id="more-information" tabindex="-1"><a class="header-anchor" href="#more-information" aria-hidden="true">#</a> More information</h3>`,4),b={href:"https://en.wikipedia.org/wiki/JSON_streaming#Line-delimited_JSON",target:"_blank",rel:"noopener noreferrer"},N=e("code",null,"json",-1),S=s('<h2 id="default-associations" tabindex="-1"><a class="header-anchor" href="#default-associations" aria-hidden="true">#</a> Default Associations</h2><ul><li><strong>Extension</strong>: <code>json-lines</code></li><li><strong>Extension</strong>: <code>jsonl</code></li><li><strong>Extension</strong>: <code>jsonlines</code></li><li><strong>Extension</strong>: <code>ldjson</code></li><li><strong>Extension</strong>: <code>murex_history</code></li><li><strong>Extension</strong>: <code>ndjson</code></li><li><strong>MIME</strong>: <code>application/json-lines</code></li><li><strong>MIME</strong>: <code>application/jsonl</code></li><li><strong>MIME</strong>: <code>application/jsonlines</code></li><li><strong>MIME</strong>: <code>application/ldjson</code></li><li><strong>MIME</strong>: <code>application/ndjson</code></li><li><strong>MIME</strong>: <code>application/x-json-lines</code></li><li><strong>MIME</strong>: <code>application/x-jsonl</code></li><li><strong>MIME</strong>: <code>application/x-jsonlines</code></li><li><strong>MIME</strong>: <code>application/x-ldjson</code></li><li><strong>MIME</strong>: <code>application/x-ndjson</code></li><li><strong>MIME</strong>: <code>text/json-lines</code></li><li><strong>MIME</strong>: <code>text/jsonl</code></li><li><strong>MIME</strong>: <code>text/jsonlines</code></li><li><strong>MIME</strong>: <code>text/ldjson</code></li><li><strong>MIME</strong>: <code>text/ndjson</code></li><li><strong>MIME</strong>: <code>text/x-json-lines</code></li><li><strong>MIME</strong>: <code>text/x-jsonl</code></li><li><strong>MIME</strong>: <code>text/x-jsonlines</code></li><li><strong>MIME</strong>: <code>text/x-ldjson</code></li><li><strong>MIME</strong>: <code>text/x-ndjson</code></li></ul><h2 id="supported-hooks" tabindex="-1"><a class="header-anchor" href="#supported-hooks" aria-hidden="true">#</a> Supported Hooks</h2><ul><li><code>Marshal()</code> Supported</li><li><code>ReadArray()</code> Works with JSON arrays. Maps are converted into arrays</li><li><code>ReadArrayWithType()</code> Works with JSON arrays. Maps are converted into arrays. Element data type is <code>json</code></li><li><code>ReadIndex()</code> Works against all properties in JSON</li><li><code>ReadMap()</code> Not currently supported.</li><li><code>ReadNotIndex()</code> Works against all properties in JSON</li><li><code>Unmarshal()</code> Supported</li><li><code>WriteArray()</code> Supported</li></ul><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>',5),O=e("code",null,"*",-1),w=e("code",null,"Marshal()",-1),k=e("code",null,"ReadArray()",-1),I=e("code",null,"ReadIndex()",-1),J=e("code",null,"[",-1),E=e("code",null,"ReadMap()",-1),v=e("code",null,"ReadNotIndex()",-1),R=e("code",null,"![",-1),A=e("code",null,"Unmarshal()",-1),C=e("code",null,"WriteArray()",-1),L=e("code",null,"[[",-1),T=e("code",null,"[",-1),W=e("code",null,"cast",-1),D=e("code",null,"csv",-1),V=e("code",null,"foreach",-1),B=e("code",null,"format",-1),U=e("code",null,"hcl",-1),H=e("code",null,"json",-1),G=e("code",null,"jsonc",-1),Y=e("code",null,"open",-1),z=e("code",null,"pretty",-1),F=e("code",null,"runtime",-1),P=e("code",null,"toml",-1),K=e("code",null,"yaml",-1);function Q(X,Z){const l=i("ExternalLinkIcon"),n=i("RouterLink");return d(),c("div",null,[h,p,m,e("p",null,[t("The following description is taken from "),e("a",_,[t("jsonlines.org"),o(l)]),t(":")]),f,g,e("p",null,[t("Example JSON lines documents taken from "),e("a",q,[t("jsonlines.org"),o(l)])]),x,e("p",null,[t("Technically the "),y,t(" Unmarshal() method supports Concatenated JSON, as described on "),e("a",j,[t("Wikipedia"),o(l)]),t(":")]),M,e("p",null,[t("This format is sometimes also referred to as LDJSON and NDJSON, as described on "),e("a",b,[t("Wikipedia"),o(l)]),t(".")]),e("p",null,[t("Murex's "),o(n,{to:"/types/json.html"},{default:a(()=>[N,t(" data-type document")]),_:1}),t(" also describes some use cases for JSON lines.")]),S,e("ul",null,[e("li",null,[o(n,{to:"/types/generic.html"},{default:a(()=>[O,t(" (generic) ")]),_:1}),t(": generic (primitive)")]),e("li",null,[o(n,{to:"/apis/Marshal.html"},{default:a(()=>[w,t(" (type)")]),_:1}),t(": Converts structured memory into a structured file format (eg for stdio)")]),e("li",null,[o(n,{to:"/apis/ReadArray.html"},{default:a(()=>[k,t(" (type)")]),_:1}),t(": Read from a data type one array element at a time")]),e("li",null,[o(n,{to:"/apis/ReadIndex.html"},{default:a(()=>[I,t(" (type)")]),_:1}),t(": Data type handler for the index, "),J,t(", builtin")]),e("li",null,[o(n,{to:"/apis/ReadMap.html"},{default:a(()=>[E,t(" (type)")]),_:1}),t(": Treat data type as a key/value structure and read its contents")]),e("li",null,[o(n,{to:"/apis/ReadNotIndex.html"},{default:a(()=>[v,t(" (type)")]),_:1}),t(": Data type handler for the bang-prefixed index, "),R,t(", builtin")]),e("li",null,[o(n,{to:"/apis/Unmarshal.html"},{default:a(()=>[A,t(" (type)")]),_:1}),t(": Converts a structured file format into structured memory")]),e("li",null,[o(n,{to:"/apis/WriteArray.html"},{default:a(()=>[C,t(" (type)")]),_:1}),t(": Write a data type, one array element at a time")]),e("li",null,[o(n,{to:"/commands/element.html"},{default:a(()=>[L,t(" (element)")]),_:1}),t(": Outputs an element from a nested structure")]),e("li",null,[o(n,{to:"/commands/index2.html"},{default:a(()=>[T,t(" (index)")]),_:1}),t(": Outputs an element from an array, map or table")]),e("li",null,[o(n,{to:"/commands/cast.html"},{default:a(()=>[W]),_:1}),t(": Alters the data type of the previous function without altering it's output")]),e("li",null,[o(n,{to:"/types/csv.html"},{default:a(()=>[D]),_:1}),t(": CSV files (and other character delimited tables)")]),e("li",null,[o(n,{to:"/commands/foreach.html"},{default:a(()=>[V]),_:1}),t(": Iterate through an array")]),e("li",null,[o(n,{to:"/commands/format.html"},{default:a(()=>[B]),_:1}),t(": Reformat one data-type into another data-type")]),e("li",null,[o(n,{to:"/types/hcl.html"},{default:a(()=>[U]),_:1}),t(": HashiCorp Configuration Language (HCL)")]),e("li",null,[o(n,{to:"/types/json.html"},{default:a(()=>[H]),_:1}),t(": JavaScript Object Notation (JSON)")]),e("li",null,[o(n,{to:"/types/jsonc.html"},{default:a(()=>[G]),_:1}),t(": Concatenated JSON")]),e("li",null,[o(n,{to:"/commands/open.html"},{default:a(()=>[Y]),_:1}),t(": Open a file with a preferred handler")]),e("li",null,[o(n,{to:"/commands/pretty.html"},{default:a(()=>[z]),_:1}),t(": Prettifies JSON to make it human readable")]),e("li",null,[o(n,{to:"/commands/runtime.html"},{default:a(()=>[F]),_:1}),t(": Returns runtime information on the internal state of Murex")]),e("li",null,[o(n,{to:"/types/toml.html"},{default:a(()=>[P]),_:1}),t(": Tom's Obvious, Minimal Language (TOML)")]),e("li",null,[o(n,{to:"/types/yaml.html"},{default:a(()=>[K]),_:1}),t(": YAML Ain't Markup Language (YAML)")]),e("li",null,[o(n,{to:"/types/mxjson.html"},{default:a(()=>[t("mxjson")]),_:1}),t(": Murex-flavoured JSON (deprecated)")])])])}const te=r(u,[["render",Q],["__file","jsonl.html.vue"]]);export{te as default};
