import{a as h}from"./index-04cccabd.js";import{d as f,a as m,q as v,b as u,o as e,c as o,g as a,e as c,w as _,t as r,E as y,p as x,i as k,_ as g}from"./index-056ef9e9.js";const I=()=>h.get("/board/v1/desc"),t=n=>(x("data-v-9ad5b82c"),n=n(),k(),n),b={class:"content"},w={class:"main-info"},C={key:0,class:"num-info"},B={key:1,class:"num-info"},S=t(()=>a("p",{class:"desc"},"腾讯Redis",-1)),E={key:0,class:"num-info"},D={key:1,class:"num-info"},M=t(()=>a("p",{class:"desc"},"阿里Redis",-1)),N={key:0,class:"num-info"},R={key:1,class:"num-info"},V=t(()=>a("p",{class:"desc"},"自建Codis",-1)),q={key:0,class:"num-info"},j={key:1,class:"num-info"},z=t(()=>a("p",{class:"desc"},"自建Cluster",-1)),A=t(()=>a("div",{class:"chart"},null,-1)),F=f({__name:"Index",setup(n){const s=m({}),p=async()=>{let i=(await I()).data;i.errorCode===0?s.value=i.data:y.error(i.msg)};return v(async()=>{await p()}),(i,G)=>{const d=u("el-button"),l=u("el-card");return e(),o("div",b,[a("div",w,[c(l,{class:"info"},{default:_(()=>[c(d,{type:"primary",icon:"el-icon-user-solid",circle:""}),s.value.txredis?(e(),o("h2",C,r(s.value.txredis),1)):(e(),o("h2",B,"0")),S]),_:1}),c(l,{class:"info"},{default:_(()=>[c(d,{type:"success",icon:"el-icon-s-data",circle:""}),s.value.aliredis?(e(),o("h2",E,r(s.value.aliredis),1)):(e(),o("h2",D,"0")),M]),_:1}),c(l,{class:"info"},{default:_(()=>[c(d,{type:"danger",icon:"el-icon-coin",circle:""}),s.value.codis?(e(),o("h2",N,r(s.value.codis),1)):(e(),o("h2",R,"0")),V]),_:1}),c(l,{class:"info"},{default:_(()=>[c(d,{type:"warning",icon:"el-icon-data-line",circle:""}),s.value.cluster?(e(),o("h2",q,r(s.value.cluster),1)):(e(),o("h2",j,"0")),z]),_:1})]),A])}}});const K=g(F,[["__scopeId","data-v-9ad5b82c"]]);export{K as default};