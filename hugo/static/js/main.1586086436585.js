!function(t){"use strict";function e(t,e){localStorage.setItem(t,JSON.stringify(e))}var n={KeyAuthentication:"settings.authentication",KeySettingsServer:"settings.server",KeySettingsInstallDefaults:"settings.install.defaults",KeyUserUuid:"app.user.uuid",KeyNotifications:"app.notifications",get:function(t,e){return localStorage.hasOwnProperty(t)?JSON.parse(localStorage.getItem(t)):e},save:e,rm:function(t){localStorage.removeItem(t)},clear:function(){localStorage.clear(),e("settings.install.defaults",!0);const t=document.querySelector('meta[name="api.server"]');e("settings.server",t?t.content:"https://learnalist.net"),e("my.edited.lists",[]),e("my.lists",[])}};function r(){}function i(t){return t()}function o(){return Object.create(null)}function s(t){t.forEach(i)}function a(t){return"function"==typeof t}function l(t,e){return t!=t?e==e:t!==e||t&&"object"==typeof t||"function"==typeof t}function c(t,e,n){t.$$.on_destroy.push(function(t,...e){if(null==t)return r;const n=t.subscribe(...e);return n.unsubscribe?()=>n.unsubscribe():n}(e,n))}function u(t,e){t.appendChild(e)}function d(t,e,n){t.insertBefore(e,n||null)}function f(t){t.parentNode.removeChild(t)}function m(t){return document.createElement(t)}function h(t){return document.createElementNS("http://www.w3.org/2000/svg",t)}function g(t){return document.createTextNode(t)}function p(){return g(" ")}function b(){return g("")}function v(t,e,n,r){return t.addEventListener(e,n,r),()=>t.removeEventListener(e,n,r)}function y(t,e,n){null==n?t.removeAttribute(e):t.getAttribute(e)!==n&&t.setAttribute(e,n)}function w(t,e){(null!=e||t.value)&&(t.value=e)}function k(t,e,n,r){t.style.setProperty(e,n,r?"important":"")}function x(t,e,n){t.classList[n?"add":"remove"](e)}let $;function S(t){$=t}const _=[],C=[],z=[],E=[],L=Promise.resolve();let M=!1;function T(t){z.push(t)}let K=!1;const N=new Set;function O(){if(!K){K=!0;do{for(let t=0;t<_.length;t+=1){const e=_[t];S(e),j(e.$$)}for(_.length=0;C.length;)C.pop()();for(let t=0;t<z.length;t+=1){const e=z[t];N.has(e)||(N.add(e),e())}z.length=0}while(_.length);for(;E.length;)E.pop()();M=!1,K=!1,N.clear()}}function j(t){if(null!==t.fragment){t.update(),s(t.before_update);const e=t.dirty;t.dirty=[-1],t.fragment&&t.fragment.p(t.ctx,e),t.after_update.forEach(T)}}const A=new Set;function P(t,e){-1===t.$$.dirty[0]&&(_.push(t),M||(M=!0,L.then(O)),t.$$.dirty.fill(0)),t.$$.dirty[e/31|0]|=1<<e%31}function H(t,e,n,l,c,u,d=[-1]){const m=$;S(t);const h=e.props||{},g=t.$$={fragment:null,ctx:null,props:u,update:r,not_equal:c,bound:o(),on_mount:[],on_destroy:[],before_update:[],after_update:[],context:new Map(m?m.$$.context:[]),callbacks:o(),dirty:d};let p=!1;if(g.ctx=n?n(t,h,(e,n,...r)=>{const i=r.length?r[0]:n;return g.ctx&&c(g.ctx[e],g.ctx[e]=i)&&(g.bound[e]&&g.bound[e](i),p&&P(t,e)),n}):[],g.update(),p=!0,s(g.before_update),g.fragment=!!l&&l(g.ctx),e.target){if(e.hydrate){const t=function(t){return Array.from(t.childNodes)}(e.target);g.fragment&&g.fragment.l(t),t.forEach(f)}else g.fragment&&g.fragment.c();e.intro&&((b=t.$$.fragment)&&b.i&&(A.delete(b),b.i(v))),function(t,e,n){const{fragment:r,on_mount:o,on_destroy:l,after_update:c}=t.$$;r&&r.m(e,n),T(()=>{const e=o.map(i).filter(a);l?l.push(...e):s(e),t.$$.on_mount=[]}),c.forEach(T)}(t,e.target,e.anchor),O()}var b,v;S(m)}let I;function R(t){let e,n,r,i,o,s,a,l,c,b;return{c(){e=m("div"),n=h("svg"),r=h("title"),i=g("info icon"),o=h("path"),a=p(),l=m("span"),c=g(t[2]),y(o,"d",s=t[4](t[1].level)),y(n,"class","w1"),y(n,"data-icon","info"),y(n,"viewBox","0 0 24 24"),k(n,"fill","currentcolor"),k(n,"width","2em"),k(n,"height","2em"),y(l,"class","lh-title ml3"),y(e,"class","flex items-center justify-center pa3 navy"),x(e,"info","info"===t[0]),x(e,"error","error"===t[0])},m(t,s,f){d(t,e,s),u(e,n),u(n,r),u(r,i),u(n,o),u(e,a),u(e,l),u(l,c),f&&b(),b=v(e,"click",D)},p(t,n){2&n&&s!==(s=t[4](t[1].level))&&y(o,"d",s),4&n&&function(t,e){e=""+e,t.data!==e&&(t.data=e)}(c,t[2]),1&n&&x(e,"info","info"===t[0]),1&n&&x(e,"error","error"===t[0])},d(t){t&&f(e),b()}}}function U(t){let e,n=t[3]&&R(t);return{c(){n&&n.c(),e=b(),this.c=r},m(t,r){n&&n.m(t,r),d(t,e,r)},p(t,[r]){t[3]?n?n.p(t,r):(n=R(t),n.c(),n.m(e.parentNode,e)):n&&(n.d(1),n=null)},i:r,o:r,d(t){n&&n.d(t),t&&f(e)}}}function D(){t.notifications.clear()}function B(e,n,r){let i;c(e,t.notifications,t=>r(1,i=t));let o,s,a;return e.$$.update=()=>{2&e.$$.dirty&&r(0,o=i.level),2&e.$$.dirty&&r(2,s=i.message),1&e.$$.dirty&&r(3,a=""!=o)},[o,i,s,a,function(t){return""==t?"":"info"==t?"M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm1 15h-2v-6h2v6zm0-8h-2V7h2v2z":"M11 15h2v2h-2zm0-8h2v6h-2zm.99-5C6.47 2 2 6.48 2 12s4.47 10 9.99 10C17.52 22 22 17.52 22 12S17.52 2 11.99 2zM12 20c-4.42 0-8-3.58-8-8s3.58-8 8-8 8 3.58 8 8-3.58 8-8 8z"}]}"function"==typeof HTMLElement&&(I=class extends HTMLElement{constructor(){super(),this.attachShadow({mode:"open"})}connectedCallback(){for(const t in this.$$.slotted)this.appendChild(this.$$.slotted[t])}attributeChangedCallback(t,e,n){this[t]=n}$destroy(){!function(t,e){const n=t.$$;null!==n.fragment&&(s(n.on_destroy),n.fragment&&n.fragment.d(e),n.on_destroy=n.fragment=null,n.ctx=[])}(this,1),this.$destroy=r}$on(t,e){const n=this.$$.callbacks[t]||(this.$$.callbacks[t]=[]);return n.push(e),()=>{const t=n.indexOf(e);-1!==t&&n.splice(t,1)}}$set(){}});function J(t){let e,n;return{c(){e=m("a"),n=g("Login"),y(e,"title","Click to login"),y(e,"href",t[0]),y(e,"class","f6 fw6 hover-red link black-70 mr2 mr3-m mr4-l dib")},m(t,r){d(t,e,r),u(e,n)},p(t,n){1&n&&y(e,"href",t[0])},d(t){t&&f(e)}}}function q(t){let e,n,i,o,s;return{c(){e=m("a"),e.textContent="Create",n=p(),i=m("a"),i.textContent="My Lists",o=p(),s=m("a"),s.textContent="Logout",y(e,"title","create, edit, share"),y(e,"href","/editor.html"),y(e,"class","f6 fw6 hover-blue link black-70 ml0 mr2-l di"),y(i,"title","Lists created by you"),y(i,"href","/lists-by-me.html"),y(i,"class","f6 fw6 hover-blue link black-70 di"),y(s,"title","Logout"),y(s,"href","/logout.html"),y(s,"class","f6 fw6 hover-blue link black-70 di ml3")},m(t,r){d(t,e,r),d(t,n,r),d(t,i,r),d(t,o,r),d(t,s,r)},p:r,d(t){t&&f(e),t&&f(n),t&&f(i),t&&f(o),t&&f(s)}}}function V(e){let n,i;function o(e,n){return null==i&&(i=!!t.loggedIn()),i?q:window.location.pathname!=e[0]?J:void 0}let s=o(e),a=s&&s(e);return{c(){n=m("div"),a&&a.c(),this.c=r,y(n,"class","fr mt0")},m(t,e){d(t,n,e),a&&a.m(n,null)},p(t,[e]){s===(s=o(t))&&a?a.p(t,e):(a&&a.d(1),a=s&&s(t),a&&(a.c(),a.m(n,null)))},i:r,o:r,d(t){t&&f(n),a&&a.d()}}}function F(t,e,n){let{loginurl:r="/login.html"}=e;return t.$set=t=>{"loginurl"in t&&n(0,r=t.loginurl)},[r]}async function G(t,e){const r={status:400,body:{}},i={username:t,password:e},o=function(){const t=n.get(n.KeySettingsServer,null);if(null===t)throw new Error("settings.server.missing");return t}()+"/api/v1/user/login",s=await fetch(o,{method:"POST",headers:{"Content-Type":"application/json"},body:JSON.stringify(i)}),a=await s.json();switch(s.status){case 200:case 403:case 400:return r.status=s.status,r.body=a,r}throw new Error("Unexpected response from the server")}function Q(t){let e,n,r,i,o,a,l,c,h,g,b,k,x,$;return{c(){e=m("form"),n=m("fieldset"),r=m("div"),i=m("label"),i.textContent="Username",o=p(),a=m("input"),l=p(),c=m("div"),h=m("label"),h.textContent="Password",g=p(),b=m("input"),k=p(),x=m("div"),x.innerHTML='<div class="w-100 items-end"><div class="fr"><div class="flex items-center mb2"><button class="db w-100" type="submit">Login</button></div> \n          <div class="flex items-center mb2"><span class="f6 link dib black">\n              or with\n              <a target="_blank" href="https://learnalist.net/api/v1/oauth/google/redirect" class="f6 link underline dib black">\n                google\n              </a></span></div></div></div>',y(i,"class","db fw6 lh-copy f6"),y(i,"for","username"),y(a,"class","pa2 input-reset ba bg-transparent b--black-20 w-100 br2"),y(a,"type","text"),y(a,"name","username"),y(a,"id","username"),y(a,"autocapitalize","none"),y(r,"class","mt3"),y(h,"class","db fw6 lh-copy f6"),y(h,"for","password"),y(b,"class","b pa2 input-reset ba bg-transparent b--black-20 w-100 br2"),y(b,"type","password"),y(b,"name","password"),y(b,"autocomplete","off"),y(b,"id","password"),y(c,"class","mv3"),y(n,"id","sign_up"),y(n,"class","ba b--transparent ph0 mh0"),y(x,"class","measure flex"),y(e,"class","measure center")},m(f,m,p){var y;d(f,e,m),u(e,n),u(n,r),u(r,i),u(r,o),u(r,a),w(a,t[0]),u(n,l),u(n,c),u(c,h),u(c,g),u(c,b),w(b,t[1]),u(e,k),u(e,x),p&&s($),$=[v(a,"input",t[4]),v(b,"input",t[5]),v(e,"submit",(y=t[2],function(t){return t.preventDefault(),y.call(this,t)}))]},p(t,e){1&e&&a.value!==t[0]&&w(a,t[0]),2&e&&b.value!==t[1]&&w(b,t[1])},d(t){t&&f(e),s($)}}}function W(t){let e;let n=Q(t);return{c(){n.c(),e=b(),this.c=r},m(t,r){n.m(t,r),d(t,e,r)},p(t,[e]){n.p(t,e)},i:r,o:r,d(t){n.d(t),t&&f(e)}}}function X(e,r,i){let o,s="",a="";return[s,a,async function(){if(""===s||""===a)return o="Please enter in a username and password",void t.notify("error",o);let e=await G(s,a);if(200!=e.status)return t.notify("error","Please try again"),void console.log(e);console.log("TODO, log them in"),n.save(n.KeyUserUuid,e.body.user_uuid),n.save(n.KeyAuthentication,e.body.token),console.log(e),t.login(e.body.token,"/")},o,function(){s=this.value,i(0,s)},function(){a=this.value,i(1,a)}]}null===n.get(n.KeySettingsInstallDefaults,null)&&n.clear(),customElements.define("login-header",class extends I{constructor(t){super(),this.shadowRoot.innerHTML="<style>a{background-color:transparent}a,div{box-sizing:border-box}.di{display:inline}.dib{display:inline-block}.fr{_display:inline}.fr{float:right}.fw6{font-weight:600}.link{text-decoration:none}.link,.link:active,.link:focus,.link:hover,.link:link,.link:visited{transition:color .15s ease-in}.link:focus{outline:1px dotted currentColor}.black-70{color:rgba(0,0,0,.7)}.hover-red:focus,.hover-red:hover{color:#ff4136}.hover-blue:focus,.hover-blue:hover{color:#357edd}.ml0{margin-left:0}.ml3{margin-left:1rem}.mr2{margin-right:.5rem}.mt0{margin-top:0}.f6{font-size:.875rem}@media screen and (min-width:30em){}@media screen and (min-width:30em) and (max-width:60em){.mr3-m{margin-right:1rem}}@media screen and (min-width:60em){.mr2-l{margin-right:.5rem}.mr4-l{margin-right:2rem}}</style>",H(this,{target:this.shadowRoot},F,V,l,{loginurl:0}),t&&(t.target&&d(t.target,this,t.anchor),t.props&&(this.$set(t.props),O()))}static get observedAttributes(){return["loginurl"]}get loginurl(){return this.$$.ctx[0]}set loginurl(t){this.$set({loginurl:t}),O()}}),customElements.define("user-login",class extends I{constructor(t){super(),this.shadowRoot.innerHTML="<style>a{background-color:transparent}button,input{font-family:inherit;font-size:100%;line-height:1.15;margin:0}button,input{overflow:visible}button{text-transform:none}button{-webkit-appearance:button}button::-moz-focus-inner{border-style:none;padding:0}button:-moz-focusring{outline:1px dotted ButtonText}fieldset{padding:.35em .75em .625em}a,div,fieldset,form,p{box-sizing:border-box}.ba{border-style:solid;border-width:1px}.b--black-20{border-color:rgba(0,0,0,.2)}.b--transparent{border-color:transparent}.br2{border-radius:.25rem}.db{display:block}.dib{display:inline-block}.flex{display:flex}.items-end{align-items:flex-end}.items-center{align-items:center}.fr{_display:inline}.fr{float:right}.b{font-weight:700}.fw6{font-weight:600}.input-reset{-webkit-appearance:none;-moz-appearance:none}.input-reset::-moz-focus-inner{border:0;padding:0}.lh-copy{line-height:1.5}.link{text-decoration:none}.link,.link:active,.link:focus,.link:hover,.link:link,.link:visited{transition:color .15s ease-in}.link:focus{outline:1px dotted currentColor}.w-100{width:100%}.black{color:#000}.bg-transparent{background-color:transparent}.pa2{padding:.5rem}.ph0{padding-left:0;padding-right:0}.mb2{margin-bottom:.5rem}.mt3{margin-top:1rem}.mv3{margin-top:1rem;margin-bottom:1rem}.mh0{margin-left:0;margin-right:0}.underline{text-decoration:underline}.f6{font-size:.875rem}.measure{max-width:30em}.center{margin-left:auto}.center{margin-right:auto}@media screen and (min-width:30em){}@media screen and (min-width:30em) and (max-width:60em){}@media screen and (min-width:60em){}</style>",H(this,{target:this.shadowRoot},X,W,l,{}),t&&t.target&&d(t.target,this,t.anchor)}}),customElements.define("notification-center",class extends I{constructor(t){super(),this.shadowRoot.innerHTML="<style>div{box-sizing:border-box}.flex{display:flex}.items-center{align-items:center}.justify-center{justify-content:center}.lh-title{line-height:1.25}.w1{width:1rem}.navy{color:#001b44}.pa3{padding:1rem}.ml3{margin-left:1rem}@media screen and (min-width:30em){}@media screen and (min-width:30em) and (max-width:60em){}@media screen and (min-width:60em){}.error{background-color:#ffdfdf}.info{background-color:#96ccff}</style>",H(this,{target:this.shadowRoot},B,U,l,{}),t&&t.target&&d(t.target,this,t.anchor)}})}(superstore);