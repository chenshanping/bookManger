<template>
    <div class="table-box">
        <div class="title">
            <h2>简单的 CRUD demo</h2>
        </div>
        <div class="query-box">
            <el-input class="query-input" v-model="queryInput" placeholder="请输入姓名搜索" @change="handleQueryName"/>
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
        <el-pagination
            background
            layout="prev, pager, next"
            :total="total"
            style="display: flex;justify-content: center;margin-top: 10px"
            v-model:current-page="curPage"
            @current-change="handleChangePage"
        />
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
     let tableForm=$ref({})
     let dialogType=$ref('add')
     let total=$ref(10)
     let curPage=$ref(1)

     /*方法*/
     /*请求table数据*/
     const getTableData= async (cur=1)=>{
         let res=await request.get('/user/list',{
             pageNum: cur,
             pageSize: 10,
         })
         tableData=res.list
         tableDataCopy=tableData
         total=res.total
         curPage=res.pageNum
     }
     getTableData()

     /*请求分页*/
     const handleChangePage=(val)=>{
         getTableData(curPage)
     }
     /*搜索*/
     let handleQueryName= async (val)=>{

         if(val.length>0){
             tableData=await request.get(`/user/list/${val}`)
             // tableData=tableData.filter(item=>(item.username).toLowerCase().match(val.toLowerCase()))
         } else{
             getTableData(curPage)
             // tableData=tableDataCopy
         }


     }
     /*编辑*/
     let handleEdit=(row)=>{
         dialogFormVisible=true
         dialogType='edit'
         tableForm={...row}
         console.log(tableForm)
     }
     /*删除第一条*/
    let handleRowClickDel= async ({ID})=>{
        console.log(ID)
        await request.delete(`/user/delete/${ID}`,)
        await getTableData(curPage)


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
     let dialogConfirm=async ()=>{
         dialogFormVisible=false
         tableForm.age=tableForm.age-0
         let index=tableForm.ID
         if( dialogType==='add'){
             /*将 number 类型转为 int类型*/
             let res=await request.post("/user",{...tableForm})

             // tableData.push({
             //     id: (tableData.length+1).toString(),
             //     ...tableForm
             // })

             await getTableData(curPage)
         }else{
             /*获取当前这条索引*/
             await request.put(`/user/update/${index}`,{...tableForm})
             await getTableData(curPage)
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
