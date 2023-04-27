import{l as z,a as A,b as F}from"./cluster-cde74a59.js";import{h as L}from"./moment-fbc5633a.js";import{d as $,a as f,r as S,q as H,b as d,s as T,o as n,c as p,g as w,e,w as l,f as V,v as q,l as G,n as O,E as i,p as W,i as j,_ as J}from"./index-94e23d1f.js";import"./index-1d9b39de.js";const K=_=>(W("data-v-88bc7fc3"),_=_(),j(),_),Q={class:"content"},X=K(()=>w("span",null,"自建Cluster集群",-1)),Z={key:0},ee={key:0},te={key:1},le={key:1},oe={key:0},ae={class:"dialog-footer"},se=$({__name:"Index",setup(_){const m=f(!1),b="100px",k=f([]),x=f([]),u=f(!1),r=S({name:"",nodes:"",password:""}),v=S({cluster_id:""}),N=(s,t)=>{const c=s[t.property];return c===void 0?"":L(c).utcOffset(8).format("YYYY-MM-DD HH:mm")},h=async()=>{let s=(await z()).data;s.errorCode===0?(m.value=!1,k.value=s.data):(m.value=!1,i.error(s.msg))},D=async s=>{if(v.cluster_id=s.ID,!v.cluster_id)i.error("未获取到的添加cluster_id");else{const t=(await A(v)).data;t.errorCode===0?x.value=t.data:i.error(t.msg),u.value=!1}},P=async()=>{if(!r.name)i.error("未获取到的添加cluster集群的名称");else if(!r.nodes)i.error("未获取到的添加cluster集群的地址");else{const s=await F(r);s.data.errorCode===0?(i.success("添加成功"),await h()):i.error(s.data.msg),u.value=!1}};return H(async()=>{m.value=!0,await h()}),(s,t)=>{const c=d("el-col"),C=d("el-button"),B=d("el-row"),E=d("el-divider"),a=d("el-table-column"),I=d("el-table"),M=d("el-card"),y=d("el-input"),g=d("el-form-item"),R=d("el-form"),U=d("el-dialog"),Y=T("loading");return n(),p("div",Q,[w("div",null,[e(M,{shadow:"never"},{default:l(()=>[e(B,null,{default:l(()=>[e(c,{span:3},{default:l(()=>[X]),_:1}),e(c,{offset:18,span:3,style:{"min-width":"120px"}},{default:l(()=>[e(C,{size:"small",type:"primary",onClick:t[0]||(t[0]=o=>u.value=!0)},{default:l(()=>[V("添加集群")]),_:1})]),_:1})]),_:1}),e(E),q((n(),G(I,{data:k.value,stripe:"",style:{width:"100%"},lazy:"",onExpandChange:D},{default:l(()=>[e(a,{label:"查看",type:"expand",width:"60"},{default:l(o=>[e(I,{data:x.value,border:!0,style:{width:"100%"},"row-key":"id",lazy:"",load:h,"tree-props":{children:"Children",hasChildren:"hasChildren"}},{default:l(()=>[e(a,{prop:"Flags",label:"身份",width:"160"}),e(a,{prop:"Address",label:"地址",width:"160"}),e(a,{prop:"LinkState",label:"LinkState",width:"160"}),e(a,{prop:"RunStatus",label:"RunStatus",width:"160"}),e(a,{prop:"SlotRange",label:"SlotRange",width:"160"}),e(a,{prop:"SlotNumber",label:"SlotNumber",width:"160"}),e(a,{prop:"CreateTime",formatter:N,label:"创建时间",width:"160"})]),_:1},8,["data"])]),_:1}),e(a,{prop:"ID",label:"ID",width:"60"}),e(a,{prop:"Name",label:"名称",width:"100"}),e(a,{prop:"Nodes",label:"内网IP"}),e(a,{prop:"Password",label:"密码",width:"120"},{default:l(o=>[o.row.NoAuth?(n(),p("div",le,[o.row.Password===""?(n(),p("span",oe,"免密登陆")):O("",!0)])):(n(),p("div",Z,[o.row.Password===""?(n(),p("span",ee,"未设置密码")):(n(),p("span",te,"已设置密码"))]))]),_:1}),e(a,{prop:"CreatedAt",formatter:N,label:"创建时间",width:"150"})]),_:1},8,["data"])),[[Y,m.value]])]),_:1})]),w("div",null,[e(U,{modelValue:u.value,"onUpdate:modelValue":t[6]||(t[6]=o=>u.value=o),title:"添加集群",width:"30%","align-center":""},{footer:l(()=>[w("span",ae,[e(C,{onClick:t[4]||(t[4]=o=>u.value=!1)},{default:l(()=>[V("取消")]),_:1}),e(C,{type:"primary",onClick:t[5]||(t[5]=o=>P())},{default:l(()=>[V("Confirm")]),_:1})])]),default:l(()=>[e(R,{model:r},{default:l(()=>[e(g,{label:"集群名称","label-width":b},{default:l(()=>[e(y,{modelValue:r.name,"onUpdate:modelValue":t[1]||(t[1]=o=>r.name=o),placeholder:"Cluster Name",autocomplete:"off"},null,8,["modelValue"])]),_:1}),e(g,{label:"集群地址","label-width":b},{default:l(()=>[e(y,{modelValue:r.nodes,"onUpdate:modelValue":t[2]||(t[2]=o=>r.nodes=o),placeholder:"127.0.0.1:6379,127.0.0.1:6380",autocomplete:"off"},null,8,["modelValue"])]),_:1}),e(g,{label:"集群密码","label-width":b},{default:l(()=>[e(y,{modelValue:r.password,"onUpdate:modelValue":t[3]||(t[3]=o=>r.password=o),type:"password",placeholder:"Cluster Password",autocomplete:"off"},null,8,["modelValue"])]),_:1})]),_:1},8,["model"])]),_:1},8,["modelValue"])])])}}});const ue=J(se,[["__scopeId","data-v-88bc7fc3"]]);export{ue as default};
