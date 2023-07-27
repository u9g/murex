import{_ as i}from"./plugin-vue_export-helper-c27b6911.js";import{r as l,o as d,c,d as e,b as t,w as n,e as o,f as r}from"./app-45f7c304.js";const s={},h=r(`<h1 id="private" tabindex="-1"><a class="header-anchor" href="#private" aria-hidden="true">#</a> <code>private</code></h1><blockquote><p>Define a private function block</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p><code>private</code> defines a function who&#39;s scope is limited to that module or source file.</p><p>Privates cannot be called from one module to another (unless they&#39;re wrapped around a global <code>function</code>) and nor can they be called from the interactive command line. The purpose of a <code>private</code> is to reduce repeated code inside a module or source file without cluttering up the global namespace.</p><h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><pre><code>private: name { code-block }
</code></pre><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><pre><code># The following cannot be entered via the command line. You need to write
# it to a file and execute it from there.

private hw {
    out &quot;Hello, World!&quot;
}

function tom {
    hw
    out &quot;My name is Tom.&quot;
}

function dick {
    hw
    out &quot;My name is Dick.&quot;
}

function harry {
    hw
    out &quot;My name is Harry.&quot;
}
</code></pre><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2><h3 id="allowed-characters" tabindex="-1"><a class="header-anchor" href="#allowed-characters" aria-hidden="true">#</a> Allowed characters</h3><p>Private names can only include any characters apart from dollar (<code>$</code>). This is to prevent functions from overwriting variables (see the order of preference below).</p><h3 id="undefining-a-private" tabindex="-1"><a class="header-anchor" href="#undefining-a-private" aria-hidden="true">#</a> Undefining a private</h3><p>Because private functions are fixed to the source file that declares them, there isn&#39;t much point in undefining them. Thus at this point in time, it is not possible to do so.</p><h3 id="order-of-preference" tabindex="-1"><a class="header-anchor" href="#order-of-preference" aria-hidden="true">#</a> Order of preference</h3><p>There is an order of precedence for which commands are looked up:</p><ol><li><p><code>runmode</code>: this is executed before the rest of the script. It is invoked by the pre-compiler forking process and is required to sit at the top of any scripts.</p></li><li><p><code>test</code> and <code>pipe</code> functions also alter the behavior of the compiler and thus are executed ahead of any scripts.</p></li><li><p>private functions - defined via <code>private</code>. Private&#39;s cannot be global and are scoped only to the module or source that defined them. For example, You cannot call a private function directly from the interactive command line (however you can force an indirect call via <code>fexec</code>).</p></li><li><p>Aliases - defined via <code>alias</code>. All aliases are global.</p></li><li><p>Murex functions - defined via <code>function</code>. All functions are global.</p></li><li><p>Variables (dollar prefixed) which are declared via <code>global</code>, <code>set</code> or <code>let</code>. Also environmental variables too, declared via <code>export</code>.</p></li><li><p>globbing: however this only applies for commands executed in the interactive shell.</p></li><li><p>Murex builtins.</p></li><li><p>External executable files</p></li></ol><p>You can override this order of precedence via the <code>fexec</code> and <code>exec</code> builtins.</p><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,19),u=e("code",null,"alias",-1),p=e("code",null,"break",-1),f=e("code",null,"exec",-1),m=e("code",null,"export",-1),_=e("code",null,"fexec",-1),b=e("code",null,"function",-1),v=e("code",null,"g",-1),x=e("code",null,"*.txt",-1),g=e("code",null,"global",-1),w=e("code",null,"let",-1),y=e("code",null,"method",-1),k=e("code",null,"set",-1),q=e("code",null,"source",-1);function D(T,A){const a=l("RouterLink");return d(),c("div",null,[h,e("ul",null,[e("li",null,[t(a,{to:"/commands/alias.html"},{default:n(()=>[u]),_:1}),o(": Create an alias for a command")]),e("li",null,[t(a,{to:"/commands/break.html"},{default:n(()=>[p]),_:1}),o(": Terminate execution of a block within your processes scope")]),e("li",null,[t(a,{to:"/commands/exec.html"},{default:n(()=>[f]),_:1}),o(": Runs an executable")]),e("li",null,[t(a,{to:"/commands/export.html"},{default:n(()=>[m]),_:1}),o(": Define an environmental variable and set it's value")]),e("li",null,[t(a,{to:"/commands/fexec.html"},{default:n(()=>[_]),_:1}),o(": Execute a command or function, bypassing the usual order of precedence.")]),e("li",null,[t(a,{to:"/commands/function.html"},{default:n(()=>[b]),_:1}),o(": Define a function block")]),e("li",null,[t(a,{to:"/commands/g.html"},{default:n(()=>[v]),_:1}),o(": Glob pattern matching for file system objects (eg "),x,o(")")]),e("li",null,[t(a,{to:"/commands/global.html"},{default:n(()=>[g]),_:1}),o(": Define a global variable and set it's value")]),e("li",null,[t(a,{to:"/commands/let.html"},{default:n(()=>[w]),_:1}),o(": Evaluate a mathematical function and assign to variable (deprecated)")]),e("li",null,[t(a,{to:"/commands/method.html"},{default:n(()=>[y]),_:1}),o(": Define a methods supported data-types")]),e("li",null,[t(a,{to:"/commands/set.html"},{default:n(()=>[k]),_:1}),o(": Define a local variable and set it's value")]),e("li",null,[t(a,{to:"/commands/source.html"},{default:n(()=>[q]),_:1}),o(": Import Murex code from another file of code block")])])])}const V=i(s,[["render",D],["__file","private.html.vue"]]);export{V as default};
