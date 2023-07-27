import{_ as o}from"./plugin-vue_export-helper-c27b6911.js";import{r as p,o as l,c as i,d as a,b as s,w as t,e as n,f as c}from"./app-45f7c304.js";const d={},r=c(`<h1 id="readindex-type-api-reference" tabindex="-1"><a class="header-anchor" href="#readindex-type-api-reference" aria-hidden="true">#</a> <code>ReadIndex()</code> (type) - API Reference</h1><blockquote><p>Data type handler for the index, <code>[</code>, builtin</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p>This is a function you would write when programming a Murex data-type.</p><p>It&#39;s called by the index, <code>[</code>, builtin.</p><p>The purpose of this function is to allow builtins to support sequential reads (where possible) and also create a standard interface for <code>[</code> (index), thus allowing it to be data-type agnostic.</p><h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><p>Registering your <code>ReadIndex()</code></p><div class="language-go line-numbers-mode" data-ext="go"><pre class="language-go"><code><span class="token comment">// To avoid data races, this should only happen inside func init()</span>
lang<span class="token punctuation">.</span>ReadIndexes<span class="token punctuation">[</span> <span class="token comment">/* your type name */</span> <span class="token punctuation">]</span> <span class="token operator">=</span> <span class="token comment">/* your readIndex func */</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><p>Example <code>ReadIndex()</code> function:</p><div class="language-go line-numbers-mode" data-ext="go"><pre class="language-go"><code><span class="token keyword">package</span> json

<span class="token keyword">import</span> <span class="token punctuation">(</span>
	<span class="token string">&quot;github.com/lmorg/murex/lang&quot;</span>
	<span class="token string">&quot;github.com/lmorg/murex/utils/json&quot;</span>
<span class="token punctuation">)</span>

<span class="token keyword">func</span> <span class="token function">index</span><span class="token punctuation">(</span>p <span class="token operator">*</span>lang<span class="token punctuation">.</span>Process<span class="token punctuation">,</span> params <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">string</span><span class="token punctuation">)</span> <span class="token builtin">error</span> <span class="token punctuation">{</span>
	<span class="token keyword">var</span> jInterface <span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span>

	b<span class="token punctuation">,</span> err <span class="token operator">:=</span> p<span class="token punctuation">.</span>Stdin<span class="token punctuation">.</span><span class="token function">ReadAll</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
	<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
		<span class="token keyword">return</span> err
	<span class="token punctuation">}</span>

	err <span class="token operator">=</span> json<span class="token punctuation">.</span><span class="token function">Unmarshal</span><span class="token punctuation">(</span>b<span class="token punctuation">,</span> <span class="token operator">&amp;</span>jInterface<span class="token punctuation">)</span>
	<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
		<span class="token keyword">return</span> err
	<span class="token punctuation">}</span>

	marshaller <span class="token operator">:=</span> <span class="token keyword">func</span><span class="token punctuation">(</span>iface <span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">)</span> <span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">,</span> <span class="token builtin">error</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
		<span class="token keyword">return</span> json<span class="token punctuation">.</span><span class="token function">Marshal</span><span class="token punctuation">(</span>iface<span class="token punctuation">,</span> p<span class="token punctuation">.</span>Stdout<span class="token punctuation">.</span><span class="token function">IsTTY</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">)</span>
	<span class="token punctuation">}</span>

	<span class="token keyword">return</span> lang<span class="token punctuation">.</span><span class="token function">IndexTemplateObject</span><span class="token punctuation">(</span>p<span class="token punctuation">,</span> params<span class="token punctuation">,</span> <span class="token operator">&amp;</span>jInterface<span class="token punctuation">,</span> marshaller<span class="token punctuation">)</span>
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2><p>While there is support for a dedicated <code>ReadNotIndex()</code> for instances of <code>![</code>, the template APIs <code>lang.IndexTemplateObject</code> and <code>lang.IndexTemplateTable</code> are both agnostic to the bang prefix.</p><h2 id="parameters" tabindex="-1"><a class="header-anchor" href="#parameters" aria-hidden="true">#</a> Parameters</h2><ol><li><code>*lang.Process</code>: Process&#39;s runtime state. Typically expressed as the variable <code>p</code></li><li><code>[]string</code>: slice of parameters used in <code>[</code></li></ol><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,17),u=a("code",null,"ReadArray()",-1),m=a("code",null,"ReadArrayWithType()",-1),k=a("code",null,"ReadNotIndex()",-1),h=a("code",null,"![",-1),b=a("code",null,"WriteArray()",-1),f=a("code",null,"[[",-1),v=a("code",null,"[",-1),x=a("code",null,"lang.IndexTemplateObject()",-1),g=a("code",null,"lang.IndexTemplateTable()",-1);function y(_,I){const e=p("RouterLink");return l(),i("div",null,[r,a("ul",null,[a("li",null,[s(e,{to:"/user-guide/bang-prefix.html"},{default:t(()=>[n("user-guide/Bang Prefix")]),_:1}),n(": Bang prefixing to reverse default actions")]),a("li",null,[s(e,{to:"/apis/ReadArray.html"},{default:t(()=>[n("apis/"),u,n(" (type)")]),_:1}),n(": Read from a data type one array element at a time")]),a("li",null,[s(e,{to:"/apis/ReadArrayWithType.html"},{default:t(()=>[n("apis/"),m,n(" (type)")]),_:1}),n(": Read from a data type one array element at a time and return the elements contents and data type")]),a("li",null,[s(e,{to:"/apis/ReadNotIndex.html"},{default:t(()=>[n("apis/"),k,n(" (type)")]),_:1}),n(": Data type handler for the bang-prefixed index, "),h,n(", builtin")]),a("li",null,[s(e,{to:"/apis/WriteArray.html"},{default:t(()=>[n("apis/"),b,n(" (type)")]),_:1}),n(": Write a data type, one array element at a time")]),a("li",null,[s(e,{to:"/commands/element.html"},{default:t(()=>[n("commands/"),f,n(" (element)")]),_:1}),n(": Outputs an element from a nested structure")]),a("li",null,[s(e,{to:"/commands/index2.html"},{default:t(()=>[n("commands/"),v,n(" (index)")]),_:1}),n(": Outputs an element from an array, map or table")]),a("li",null,[s(e,{to:"/apis/lang.IndexTemplateObject.html"},{default:t(()=>[n("apis/"),x,n(" (template API)")]),_:1}),n(": Returns element(s) from a data structure")]),a("li",null,[s(e,{to:"/apis/lang.IndexTemplateTable.html"},{default:t(()=>[n("apis/"),g,n(" (template API)")]),_:1}),n(": Returns element(s) from a table")])])])}const T=o(d,[["render",y],["__file","ReadIndex.html.vue"]]);export{T as default};
