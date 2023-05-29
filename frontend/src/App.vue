
<script setup>
import { Edit, Plus, Delete, Search, Refresh, Upload } from '@element-plus/icons-vue'

import { ref ,reactive} from 'vue'

import axios from 'axios'

import { makePlaceString, makeClassString, makeDateString } from './components/ResolveId.js'

import { lab_places } from './components/labPlaces.js'
import { classes } from './components/classes.js'

import { ElMessageBox } from 'element-plus'




const items = ref([])
const filteredItems = ref([])
const searchQuery = ref('')

// new value
const newItem = ref({
    name: "",
    id: "",
    place: "",
    place_row: null,//num
    place_col: null,//num
    owner: "AiLab",//string
    class: ""//array(2)
})

const members = ['AiLab', 'lsm', 'gdx', 'chb']
const newItemDialogVisible = ref(false)

function updateItem() {
    const url = "http://localhost:8080/data/getall"
    axios.get(url)
        .then(res => {
            items.value = res.data

            for (var ii in res.data) {
                items.value[ii]['place'] = makePlaceString(items.value[ii]['id'])
                items.value[ii]['class'] = makeClassString(items.value[ii]['id'])

            }
            filteredItems.value = items.value
            // console.log(items.value)
        })
        .catch(error => {
            ElMessageBox.alert(error)
        })

}

function deleteItem(index) {
    const delete_id = items.value[index].id
    console.log(delete_id)

    ElMessageBox
        .confirm(`确定删除${items.value[index].owner}的物品:"${items.value[index].name}"?`)
        .then((action) => {
            if (action == "confirm") {
                const url = "http://localhost:8080/data/delete?id=" + delete_id
                axios.get(url)
                    .then(res => {
                        console.log(res.data)


                        items.value = items.value.filter(item =>
                            item['id'] != res.data
                        )

                        filteredItems.value = items.value
                    })
                    .catch(error => {
                        ElMessageBox.alert(error)
                    })
            }
        })

}

function addItem() {
    const url = "http://localhost:8080/data/add"
    console.log(newItem.value)

    
    const nv = newItem.value;

    if (!nv.name || !nv.price || !nv.owner || !nv.place ||!nv.class) {
        ElMessageBox.alert("unfilled")
        return
    }
    const newItemReal = {
        id: nv.place[0] + nv.place[1] + (nv.place_row ? 0 : 1) + (nv.place_col ? 0 : 1) + nv.class[0] + nv.class[1] + '00',
        name: nv.name,
        price: Number(nv.price),
        owner: nv.owner,
        time: makeDateString()
    }

    console.log(newItemReal)
    if (nv) {
        axios({
            url: url,
            params: newItemReal.value
        })
        axios.get(url)
            .then(res => {

            })
            .catch(error => {
                ElMessageBox.alert(error)
            })
    }
    else {
        ElMessageBox.alert('unfill')
    }
    updateItem()
}


function onSearch() {
    filteredItems.value = items.value.filter(function (item) {
        const searchQueryLower = searchQuery.value.toLowerCase()

        return false ||
            item['id'].includes(searchQueryLower) ||
            item['place'].includes(searchQueryLower) ||
            item['price'].toString().includes(searchQueryLower) ||
            item['time'].includes(searchQueryLower) ||
            item['owner'].toLowerCase().includes(searchQueryLower) ||
            item['name'].toLowerCase().includes(searchQueryLower)
        // item['class'].toLowerCase().includes(searchQueryLower) 
    })
}


updateItem()

filteredItems.value = items.value
</script>
<template>
    <!-- 添加物品dialog -->
    <el-dialog v-model="newItemDialogVisible" title="添加物品" style="width: max-content;">
        <el-form :model="form" label-width="5rem">
            <el-form-item label="物品名称" >
                <el-input v-model="newItem.name" clearable />
            </el-form-item>
            <el-form-item label="物品价值" >
                <el-input v-model="newItem.price" clearable />
            </el-form-item>

            <el-form-item label="所者">
                <el-select v-model="newItem.owner" class="m-2" placeholder="AiLab">
                    <el-option v-for="mem in members" :label="mem" :value="mem" />
                </el-select>
            </el-form-item>

            <el-form-item label="位置">
                <el-cascader v-model="newItem.place" :options="lab_places" />
                <el-input-number v-model="newItem.place_row" min="1" :disabled="(newItem.place[1] === '0')" />行
                <el-input-number v-model="newItem.place_col" min="1" :disabled="(newItem.place[1] === '0')" />列
            </el-form-item>
            <el-form-item label="类型">
                <el-cascader v-model="newItem.class" :options="classes" />

            </el-form-item>

            <el-button type="primary" @click="addItem">
                Upload<el-icon class="el-icon--right">
                    <Upload />
                </el-icon>
            </el-button>

        </el-form>

    </el-dialog>

    <div class="common-layout">
        <el-container>
            <!-- <el-header>
                    
            </el-header> -->
            <el-main>
                <el-button type="primary" :icon="Plus" @click="newItemDialogVisible = true;" />

                <el-button type="primary" :icon="Refresh" @click="updateItem" />

                <el-input v-model="searchQuery" placeholder="请输入搜索关键字" :prefix-icon="Search" @input="onSearch">
                </el-input>


                <el-table :data="filteredItems" stripe style="width: 100%">
                    <el-table-column fixed prop="name" label="名称" />
                    <!-- <el-table-column prop="id" label="Id" /> -->
                    <el-table-column prop="place" label="位置" width="180rem" />
                    <el-table-column prop="class" label="类型" width="140rem" />

                    <el-table-column prop="owner" label="所有者" />
                    <el-table-column prop="price" label="价值" />
                    <el-table-column prop="time" label="购入时间" width="100rem" />

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