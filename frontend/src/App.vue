<template>
    <div class="table-box">
        <div class="title">
            <h2>简单的CRUDdemo</h2>
        </div>
        <div class="query-box">
            <el-input class="query-input" v-model="queryInput" placeholder="请输入姓名搜索" />
            <el-button  text type="primary" @click="handleAdd">增加</el-button>
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
            <el-table-column fixed prop="id" label="ID" width="120" />
            <el-table-column prop="name" label="姓名" width="120" />
            <el-table-column prop="sex" label="性别" width="120" />
            <el-table-column prop="age" label="年龄" width="120" />
            <el-table-column prop="phone" label="手机号码" width="250" />
            <el-table-column fixed="right" label="操作" width="120">
                <template #default>
                    <el-button link type="primary" size="small" @click="handleRowClick">
                        删除
                    </el-button>
                    <el-button link type="primary" size="small">
                        编辑
                    </el-button>
                </template>
            </el-table-column>
        </el-table>
       <!-- dialog       -->
        <el-dialog v-model="dialogFormVisible" title="Shipping address" >
<!--            :label-position="labelPosition" :label-width="300"-->
            <el-form :model="tableForm" >
                <el-form-item label="id" :label-width="100">
                    <el-input v-model="tableForm.id" autocomplete="off" />
                 </el-form-item>
                <el-form-item label="姓名" :label-width="100">
                    <el-input v-model="tableForm.name" autocomplete="off" />
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
                <el-button type="primary" @click="dialogFormVisible = false">
                    确认
                </el-button>
            </span>
            </template>
        </el-dialog>
    </div>
</template>
<script setup>
     import {ref} from "vue";
     /*数据*/
     let queryInput=ref("")
     let tableData=ref([
         {
         id: '1',
         name: '陈善平',
         sex: '男',
         age: 18,
        phone:"18824682674"
     },
         {
             id: '2',
             name: '张三',
             sex: '男',
             age: 18,
             phone:"18824682674"
         },
         {
             id: '3',
             name: '王四',
             sex: '男',
             age: 18,
             phone:"18824682674"
         },
         {
             id: '4',
             name: '吴五',
             sex: '男',
             age: 18,
             phone:"18824682674"
         },
     ])
     let multipleSelection=ref([])
     let dialogFormVisible=ref(false)
     let tableForm=ref({
         id:1,
         name:'张三',
         sex:'男',
         age:12,
         phone:"18824682674"
     })
     /*方法*/
    let handleRowClick=()=>{
        console.log('click')
    }
     let handleSelectionChange = (val) => {
         multipleSelection.value = val
         console.log(val)
     }
     let handleAdd=()=>{
         dialogFormVisible.value=true
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
    width: 300px;
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
