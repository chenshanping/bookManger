<template>
    <div class="table-box">
        <div class="title">
            <h2>简单的 CRUD demo</h2>
        </div>
        <div class="query-box">
            <el-input class="query-input" v-model="queryInput" placeholder="请输入姓名搜索" @input="handleQueryName"/>
            <div class="btn-list">
                <el-button  text type="primary" @click="handleAdd">增加</el-button>
                <el-button  text type="danger" @click="handleDelList" v-if="multipleSelection.length>0">删除</el-button>
            </div>

        </div>
        <!--table-->
        <el-table
            border
            ref="multipleTableRef"
            :data="tableData"
            style="width: 100%"
            @selection-change="handleSelectionChange"
        >
            <el-table-column type="selection" width="55" />
            <el-table-column prop="username" label="姓名" width="120" />
            <el-table-column prop="sex" label="性别" width="120" />
            <el-table-column prop="age" label="年龄" width="120" />
            <el-table-column prop="phone" label="手机号码" width="250" />
            <el-table-column fixed="right" label="操作" width="120">
                <template #default="scope">
                    <el-button link type="primary" size="small" @click="handleRowClickDel(scope.row)"
                    style="color: #ff1a44"
                    >
                        删除
                    </el-button>
                    <el-button link type="primary" size="small" @click="handleEdit(scope.row)">
                        编辑
                    </el-button>
                </template>
            </el-table-column>
        </el-table>
       <!-- dialog       -->
        <el-dialog v-model="dialogFormVisible" :title="dialogType ==='add' ? '新增':'编辑'" >
<!--            :label-position="labelPosition" :label-width="300"-->
            <el-form :model="tableForm" >
                <el-form-item label="姓名" :label-width="100">
                    <el-input v-model="tableForm.username" autocomplete="off" />
                </el-form-item>
                <el-form-item label="性别" :label-width="100">
                    <el-input v-model="tableForm.sex" autocomplete="off" />
                </el-form-item>
                <el-form-item label="年龄" :label-width="100">
                    <el-input v-model="tableForm.age" autocomplete="off" />
                </el-form-item>
                <el-form-item label="手机号码" :label-width="100">
                    <el-input v-model="tableForm.phone" autocomplete="off" />
                </el-form-item>
            </el-form>
            <template #footer>
            <span class="dialog-footer">
                <el-button type="primary" @click="dialogConfirm">
                    确认
                </el-button>
            </span>
            </template>
        </el-dialog>
    </div>
</template>
<script setup>
     import {ref} from "vue";
     import axios from "axios";
     import request from "./utils/request.js";
     /*数据*/
     let queryInput=$ref("")
     let tableData=$ref([])
     let tableDataCopy
     let multipleSelection=$ref([])
     let dialogFormVisible=$ref(false)
     let tableForm=$ref({
         id:1,
         username:'张三',
         sex:'男',
         age:12,
         phone:"18824682674"
     })
     let dialogType=$ref('add')
     /*方法*/

     const getTableData= async (cur=1)=>{
         let res=await request.get('/list',{
             pageNum: -1,
             pageSize: -1,
         })
         tableData=res.list
         tableDataCopy=tableData
         console.log(res.list)
     }
     getTableData()

     /*搜索*/
     let handleQueryName=(val)=>{
         if(val.length>0){
             tableData=tableData.filter(item=>(item.username).toLowerCase().match(val.toLowerCase()))
         } else{
             tableData=tableDataCopy
         }


     }
     /*编辑*/
     let handleEdit=(row)=>{
         dialogFormVisible=true
         dialogType='edit'
         tableForm={...row}
     }
     /*删除第一条*/
    let handleRowClickDel= async ({ID})=>{
        let index=tableData.findIndex(item=>item.ID===ID)
        tableData.splice(index,1)
        let res=await request.delete( `/delete/${ID}`)

    }
    /*删除多条*/
     let handleDelList = ()=>{
        multipleSelection.forEach(ID=>{
            handleRowClickDel({ID})

        })

     }
     /*多选*/
     let handleSelectionChange = (val) => {

         // multipleSelection = val
         // console.log(val)
         multipleSelection=[]
         val.forEach(
             item=>{
                 multipleSelection.push(item.ID)
             }
         )
     }
    /*新增*/
     let handleAdd=()=>{
         dialogFormVisible=true
         tableForm={}
         dialogType='add'
     }
     let dialogConfirm=()=>{
         dialogFormVisible=false
         if( dialogType==='add'){
             tableData.push({
                 id: (tableData.length+1).toString(),
                 ...tableForm
             })
         }else{
             /*获取当前这条索引*/
             let index=tableData.findIndex(item=>item.id===tableForm.id)
             tableData[index]=tableForm
         }

     }

</script>
<style scoped>
.table-box {
    margin: 200px auto;
    width: 800px;

}
 .title{
     text-align: center;

 }
 .query-box{
     display: flex;
     justify-content: space-between;
     margin-bottom: 20px;
 }
 .query-input{
    width: 200px;
 }
 .el-dialog{
     margin: 200px auto;
 }

.el-button--text {
    margin-right: 15px;
}
.dialog-footer button:first-child {
    margin-right: 10px;
}

</style>
