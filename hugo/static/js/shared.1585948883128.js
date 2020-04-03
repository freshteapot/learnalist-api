var superstore=function(e){"use strict";var t=function(e,t){return e(t={exports:{}},t.exports),t.exports}((function(e,t){var n;n=function(){function e(){for(var e=0,t={};e<arguments.length;e++){var n=arguments[e];for(var o in n)t[o]=n[o]}return t}function t(e){return e.replace(/(%[0-9A-Z]{2})+/g,decodeURIComponent)}return function n(o){function r(){}function i(t,n,i){if("undefined"!=typeof document){"number"==typeof(i=e({path:"/"},r.defaults,i)).expires&&(i.expires=new Date(1*new Date+864e5*i.expires)),i.expires=i.expires?i.expires.toUTCString():"";try{var c=JSON.stringify(n);/^[\{\[]/.test(c)&&(n=c)}catch(e){}n=o.write?o.write(n,t):encodeURIComponent(String(n)).replace(/%(23|24|26|2B|3A|3C|3E|3D|2F|3F|40|5B|5D|5E|60|7B|7D|7C)/g,decodeURIComponent),t=encodeURIComponent(String(t)).replace(/%(23|24|26|2B|5E|60|7C)/g,decodeURIComponent).replace(/[\(\)]/g,escape);var s="";for(var a in i)i[a]&&(s+="; "+a,!0!==i[a]&&(s+="="+i[a].split(";")[0]));return document.cookie=t+"="+n+s}}function c(e,n){if("undefined"!=typeof document){for(var r={},i=document.cookie?document.cookie.split("; "):[],c=0;c<i.length;c++){var s=i[c].split("="),a=s.slice(1).join("=");n||'"'!==a.charAt(0)||(a=a.slice(1,-1));try{var l=t(s[0]);if(a=(o.read||o)(a,l)||t(a),n)try{a=JSON.parse(a)}catch(e){}if(r[l]=a,e===l)break}catch(e){}}return e?r[e]:r}}return r.set=i,r.get=function(e){return c(e,!1)},r.getJSON=function(e){return c(e,!0)},r.remove=function(t,n){i(t,"",e(n,{expires:-1}))},r.defaults={},r.withConverter=n,r}((function(){}))},e.exports=n()}));function n(){}const o=[];function r(e,t=n){let r;const i=[];function c(t){if(c=t,((n=e)!=n?c==c:n!==c||n&&"object"==typeof n||"function"==typeof n)&&(e=t,r)){const t=!o.length;for(let t=0;t<i.length;t+=1){const n=i[t];n[1](),o.push(n,e)}if(t){for(let e=0;e<o.length;e+=2)o[e][0](o[e+1]);o.length=0}}var n,c}return{set:c,update:function(t){c(t(e))},subscribe:function(o,s=n){const a=[o,s];return i.push(a),1===i.length&&(r=t(c)||n),o(e),()=>{const e=i.indexOf(a);-1!==e&&i.splice(e,1),0===i.length&&(r(),r=null)}}}}function i(e,t){localStorage.setItem(e,JSON.stringify(t))}var c={KeyAuthentication:"settings.authentication",KeySettingsServer:"settings.server",KeySettingsInstallDefaults:"settings.install.defaults",KeyUserUuid:"app.user.uuid",KeyNotifications:"app.notifications",get:function(e,t){return localStorage.hasOwnProperty(e)?JSON.parse(localStorage.getItem(e)):t},save:i,rm:function(e){localStorage.removeItem(e)},clear:function(){localStorage.clear(),i("settings.install.defaults",!0);const e=document.querySelector('meta[name="api.server"]');i("settings.server",e?e.content:"https://learnalist.net"),i("my.edited.lists",[]),i("my.lists",[])}};const s={level:"",message:""},a=JSON.parse(JSON.stringify(s));let l=JSON.parse(JSON.stringify(s));const u=c.get(c.KeyNotifications,null);null!==u&&(l=u);const{subscribe:f,update:g,set:p}=r(l);const d={subscribe:f,add:(e,t)=>{g(n=>(n.level=e,n.message=t,c.save(c.KeyNotifications,n),n))},clear:()=>{c.rm(c.KeyNotifications),p(a)}},m=r(0);return e.count=m,e.loggedIn=()=>(console.log("am I logged in"),console.log("Should I check local storage"),!!t.get("x-authentication-bearer")),e.login=(e,n)=>{t.set("x-authentication-bearer",e),n||(n="/welcome.html"),window.location=n},e.logout=e=>{console.log("I want to be logged out."),t.remove("x-authentication-bearer"),localStorage.clear(),"#"!==e&&(e||(e="/welcome.html"),window.location=e)},e.notifications=d,e.notify=(e,t)=>{d.add(e,t)},e}({});
