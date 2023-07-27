import{_ as n}from"./plugin-vue_export-helper-c27b6911.js";import{r as i,o as l,c as r,d as e,b as o,w as s,e as t,f as d}from"./app-45f7c304.js";const u={},c=d(`<h1 id="suffix" tabindex="-1"><a class="header-anchor" href="#suffix" aria-hidden="true">#</a> <code>suffix</code></h1><blockquote><p>Prefix a string to every item in a list</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p>Takes a list from STDIN and returns that same list with each element suffixed.</p><h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><pre><code>\`&lt;stdin&gt;\` -&gt; suffix str -&gt; \`&lt;stdout&gt;\`
</code></pre><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><pre><code>» ja: [Monday..Wednesday] -&gt; suffix foobar
[
    &quot;Mondayfoobar&quot;,
    &quot;Tuesdayfoobar&quot;,
    &quot;Wednesdayfoobar&quot;
]
</code></pre><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2><p>Supported data types can queried via <code>runtime</code></p><pre><code>runtime: --marshallers
runtime: --unmarshallers
</code></pre><h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2><ul><li><code>suffix</code></li><li><code>list.suffix</code></li></ul><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,14),h=e("code",null,"a",-1),m=e("code",null,"count",-1),f=e("code",null,"ja",-1),_=e("code",null,"lang.MarshalData()",-1),p=e("code",null,"lang.UnmarshalData()",-1),x=e("code",null,"left",-1),y=e("code",null,"prefix",-1),g=e("code",null,"right",-1),b=e("code",null,"runtime",-1);function v(k,q){const a=i("RouterLink");return l(),r("div",null,[c,e("ul",null,[e("li",null,[o(a,{to:"/commands/a.html"},{default:s(()=>[h,t(" (mkarray)")]),_:1}),t(": A sophisticated yet simple way to build an array or list")]),e("li",null,[o(a,{to:"/commands/count.html"},{default:s(()=>[m]),_:1}),t(": Count items in a map, list or array")]),e("li",null,[o(a,{to:"/commands/ja.html"},{default:s(()=>[f,t(" (mkarray)")]),_:1}),t(": A sophisticated yet simply way to build a JSON array")]),e("li",null,[o(a,{to:"/apis/lang.MarshalData.html"},{default:s(()=>[_,t(" (system API)")]),_:1}),t(": Converts structured memory into a Murex data-type (eg for stdio)")]),e("li",null,[o(a,{to:"/apis/lang.UnmarshalData.html"},{default:s(()=>[p,t(" (system API)")]),_:1}),t(": Converts a Murex data-type into structured memory")]),e("li",null,[o(a,{to:"/commands/left.html"},{default:s(()=>[x]),_:1}),t(": Left substring every item in a list")]),e("li",null,[o(a,{to:"/commands/prefix.html"},{default:s(()=>[y]),_:1}),t(": Prefix a string to every item in a list")]),e("li",null,[o(a,{to:"/commands/right.html"},{default:s(()=>[g]),_:1}),t(": Right substring every item in a list")]),e("li",null,[o(a,{to:"/commands/runtime.html"},{default:s(()=>[b]),_:1}),t(": Returns runtime information on the internal state of Murex")])])])}const N=n(u,[["render",v],["__file","suffix.html.vue"]]);export{N as default};
