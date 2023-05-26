
<script setup>
import { Edit, Plus, Delete, Search, Refresh } from '@element-plus/icons-vue'

import { ref } from 'vue'

import axios from 'axios'

import { makePlaceString, makeClassString } from './components/ResolveId.js'


import Topbar from './components/Topbar.vue'


const items = ref([])
const searchQuery = ref('')

// new value
const nvName = ref()
const nvPlace = ref()
const nvOwner = ref()
const nvTime = ref()
const nvClass = ref()
const nvPrice = ref()
function updateItem() {
    const url = "http://localhost:8080/data/getall"
    axios.get(url)
        .then(res => {
            items.value = res.data

            for (var ii in res.data) {
                items.value[ii]['place'] = makePlaceString(items.value[ii]['id'])
                items.value[ii]['class'] = makeClassString(items.value[ii]['id'])

            }

            console.log(items.value)
        })
        .catch(error => {
            alert(error)
        })

}

function deleteItem(index) {
    const delete_id = items.value[index].id
    console.log(delete_id)

    if (confirm(`确定删除${items.value[index].owner}的物品:"${items.value[index].name}"?`) == false) {
        return
    }
    const url = "http://localhost:8080/data/delete?id=" + delete_id
    axios.get(url)
        .then(res => {
            console.log(res.data)

            items.value = items.value.filter(item =>
                item['id'] != res.data
            )


        })
        .catch(error => {
            alert(error)
        })
}

function addItem() {
    const url = "http://localhost:8080/data/add"

    if (newItem.value) {
        axios({
            url: url,
            params: newItem.value
        })
        axios.get(url)
            .then(res => {

            })
            .catch(error => {
                alert(error)
            })
    }
    else {
        alert("unfill")
    }
}

function searchItem() {

}
</script>
<template>
    <div class="common-layout">
        <el-container>
            <!-- <el-header>
                    
            </el-header> -->
            <el-main>
                <el-button type="primary" :icon="Plus" @click="addItem" />

                <el-button type="primary" :icon="Refresh" @click="updateItem" />

                <el-input v-model="searchQuery" placeholder="请输入搜索关键字" class="input-with-select">
                    <template #append>
                        <el-button :icon="Search" @Click="searchItem" />
                    </template>
                </el-input>

                <el-button type="primary" :icon="Search" />




                <el-table :data="items" stripe style="width: 100%">
                    <el-table-column fixed prop="name" label="名称" />
                    <!-- <el-table-column prop="id" label="Id" /> -->
                    <el-table-column prop="place" label="位置" width="130rem" />
                    <el-table-column prop="class" label="类型" />

                    <el-table-column prop="owner" label="所有者" />
                    <el-table-column prop="price" label="价值" />
                    <el-table-column prop="time" label="购入时间" />

                    <el-table-column fixed="right" label="操作">
                        <template #default="scope">
                            <el-button link type="primary" :icon="Delete" @click.prevent="deleteItem(scope.$index)" />
                            <el-button link type="primary" :icon="Edit" />
                        </template>
                    </el-table-column>
                </el-table>
            </el-main>
        </el-container>
    </div>
</template>
<style></style>