import{l as R,a as T,b as W,o as j}from"./codis-7a3bd4ff.js";import{d as A,a as C,r as F,q as H,b as n,o as p,c as v,g as _,e as l,w as a,F as U,k as h,l as M,f as m,t as J,n as S,E as r,p as K,i as P,_ as Q}from"./index-056ef9e9.js";import"./index-04cccabd.js";const X=y=>(K("data-v-c36ae760"),y=y(),P(),y),Y={class:"content"},Z=X(()=>_("span",null,"自建codis集群",-1)),ee=["src"],le={class:"dialog-footer"},oe={key:0},ae={key:1},te={class:"dialog-footer"},de=A({__name:"Index",setup(y){const I=C([]),$=C([]),V=C(!1),b=C(!1),u="100px",E=[{value:"dilatation",label:"扩容"},{value:"shrinkage",label:"缩容"}],t=F({curl:"",cluster_name:"",add_proxy:"",add_server:"",del_proxy:0,del_group:0,op_type:""}),d=F({curl:"",cname:""}),w=async()=>{let s=(await R()).data;s.errorCode===0?I.value=s.data.lists:r.error(s.msg)},L=async()=>{const s=await T(d);s.data.errorCode===0?($.value=s.data.data,await w()):r.error(s.data.msg)},q=async()=>{if(console.log(d),!d.curl)r.error("未获取到的新增codis平台地址");else if(!d.cname)r.error("未获取到的新增codis平台名称");else{const s=await W(d);s.data.errorCode===0?(r.success("添加成功"),await w()):r.error(s.data.msg),V.value=!1}},D=async()=>{if(t.curl=d.curl,!t.cluster_name||!t.curl)r.error("未获取到集群名和集群地址");else{console.log(t);const s=await j(t);s.data.errorCode===0?(r.success(s.data.data),await w()):(r.error(s.data.msg),r.error(s.data.data)),b.value=!1}};return H(async()=>{await w()}),(s,o)=>{const g=n("el-col"),k=n("el-option"),x=n("el-select"),G=n("el-link"),f=n("el-button"),O=n("el-row"),N=n("el-card"),c=n("el-input"),i=n("el-form-item"),z=n("el-form"),B=n("el-dialog");return p(),v("div",Y,[_("div",null,[l(N,{shadow:"never"},{default:a(()=>[l(O,null,{default:a(()=>[l(g,{span:3},{default:a(()=>[Z]),_:1}),l(g,{span:2},{default:a(()=>[l(x,{modelValue:d.curl,"onUpdate:modelValue":o[0]||(o[0]=e=>d.curl=e),placeholder:"请选择平台地址",onChange:o[1]||(o[1]=e=>L())},{default:a(()=>[(p(!0),v(U,null,h(I.value,e=>(p(),M(k,{key:e.Cname,label:e.Cname,value:e.Curl},null,8,["label","value"]))),128))]),_:1},8,["modelValue"])]),_:1}),l(g,{offset:1,span:10},{default:a(()=>[l(G,{href:d.curl,target:"_blank",type:"success"},{default:a(()=>[m(J(d.curl),1)]),_:1},8,["href"])]),_:1}),l(g,{offset:2,span:2,style:{"min-width":"120px"}},{default:a(()=>[l(f,{size:"small",onClick:o[2]||(o[2]=e=>b.value=!0)},{default:a(()=>[m("扩缩容")]),_:1})]),_:1}),l(g,{span:3,style:{"min-width":"120px"}},{default:a(()=>[l(f,{size:"small",type:"primary",onClick:o[3]||(o[3]=e=>V.value=!0)},{default:a(()=>[m("添加codis平台")]),_:1})]),_:1})]),_:1})]),_:1})]),_("div",null,[l(N,{shadow:"never"},{default:a(()=>[_("iframe",{src:d.curl,frameborder:"0",width:"100%",height:"100%",class:"codis_dashboard"},null,8,ee)]),_:1})]),_("div",null,[l(B,{modelValue:V.value,"onUpdate:modelValue":o[8]||(o[8]=e=>V.value=e),title:"新增codis平台地址",width:"30%","align-center":""},{footer:a(()=>[_("span",le,[l(f,{onClick:o[6]||(o[6]=e=>V.value=!1)},{default:a(()=>[m("取消")]),_:1}),l(f,{type:"primary",onClick:o[7]||(o[7]=e=>q())},{default:a(()=>[m("确定")]),_:1})])]),default:a(()=>[l(z,{model:d},{default:a(()=>[l(i,{label:"平台名称","label-width":u},{default:a(()=>[l(c,{modelValue:d.cname,"onUpdate:modelValue":o[4]||(o[4]=e=>d.cname=e),autocomplete:"off"},null,8,["modelValue"])]),_:1}),l(i,{label:"平台地址","label-width":u},{default:a(()=>[l(c,{modelValue:d.curl,"onUpdate:modelValue":o[5]||(o[5]=e=>d.curl=e),autocomplete:"off"},null,8,["modelValue"])]),_:1})]),_:1},8,["model"])]),_:1},8,["modelValue"]),l(B,{modelValue:b.value,"onUpdate:modelValue":o[17]||(o[17]=e=>b.value=e),title:"扩缩容codis集群",width:"30%","align-center":""},{footer:a(()=>[_("span",te,[l(f,{onClick:o[15]||(o[15]=e=>b.value=!1)},{default:a(()=>[m("取消")]),_:1}),l(f,{type:"primary",onClick:o[16]||(o[16]=e=>D())},{default:a(()=>[m("确定")]),_:1})])]),default:a(()=>[l(z,{model:d},{default:a(()=>[l(i,{label:"操作类型","label-width":u},{default:a(()=>[l(x,{modelValue:t.op_type,"onUpdate:modelValue":o[9]||(o[9]=e=>t.op_type=e),class:"m-2",placeholder:"操作",size:"large"},{default:a(()=>[(p(),v(U,null,h(E,e=>l(k,{key:e.value,label:e.label,value:e.value},null,8,["label","value"])),64))]),_:1},8,["modelValue"])]),_:1}),l(i,{label:"选择集群","label-width":u},{default:a(()=>[l(x,{modelValue:t.cluster_name,"onUpdate:modelValue":o[10]||(o[10]=e=>t.cluster_name=e),class:"m-2",placeholder:"集群",size:"large"},{default:a(()=>[(p(!0),v(U,null,h($.value,e=>(p(),M(k,{key:e,label:e,value:e},null,8,["label","value"]))),128))]),_:1},8,["modelValue"])]),_:1}),t.op_type==="dilatation"?(p(),v("div",oe,[l(i,{label:"proxy列表","label-width":u},{default:a(()=>[l(c,{modelValue:t.add_proxy,"onUpdate:modelValue":o[11]||(o[11]=e=>t.add_proxy=e),placeholder:"127.0.0.1:11081,127.0.0.1:11082",autocomplete:"off"},null,8,["modelValue"])]),_:1}),l(i,{label:"Redis列表","label-width":u},{default:a(()=>[l(c,{modelValue:t.add_server,"onUpdate:modelValue":o[12]||(o[12]=e=>t.add_server=e),placeholder:"127.0.0.1:6379,127.0.0.1:6380",autocomplete:"off"},null,8,["modelValue"])]),_:1})])):S("",!0),t.op_type==="shrinkage"?(p(),v("div",ae,[l(i,{label:"proxy数量","label-width":u},{default:a(()=>[l(c,{modelValue:t.del_proxy,"onUpdate:modelValue":o[13]||(o[13]=e=>t.del_proxy=e),modelModifiers:{number:!0},placeholder:"1",autocomplete:"off"},null,8,["modelValue"])]),_:1}),l(i,{label:"group数量","label-width":u},{default:a(()=>[l(c,{modelValue:t.del_group,"onUpdate:modelValue":o[14]||(o[14]=e=>t.del_group=e),modelModifiers:{number:!0},placeholder:"1",autocomplete:"off"},null,8,["modelValue"])]),_:1})])):S("",!0)]),_:1},8,["model"])]),_:1},8,["modelValue"])])])}}});const ue=Q(de,[["__scopeId","data-v-c36ae760"]]);export{ue as default};