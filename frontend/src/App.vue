<script setup>

import Topbar from './components/Topbar.vue'
// import AddItemDialog from './components/AddItemDialog.vue'

import axios from 'axios'
import { ref } from 'vue'



let items = ref([])
const filteredItems = ref([])
const searchQuery = ref('')

function filterItems() {

    filteredItems.value = items.value.filter(function(item)
    {
        const searchQueryLower = searchQuery.value.toLowerCase()

        return false || 
        item['owner'].toLowerCase().includes(searchQueryLower) ||
        item['id'].toLowerCase().includes(searchQueryLower) ||
        item['price'].toString().includes(searchQueryLower) ||
        item['time'].toLowerCase().includes(searchQueryLower) ||
        item['name'].toLowerCase().includes(searchQueryLower) 
    }
    )

    // console.log(filteredItems.value)
    console.log(searchQuery.value, filteredItems.value)


}

function updateItem() {
    const url = "http://localhost:8080/data/get"
    axios.get(url)
        .then(res => {
            items.value = res.data
            console.log(items)
        })
        .catch(error => {
            alert(error)
        })
    
    filteredItems.value = items.value
}


</script>

<template>
    <Topbar></Topbar>
    <div id="container">

        <form id="search">
            Search <input id="query_box" name="query" @input="filterItems" v-model="searchQuery">
        </form>

        <div id="btns">
            <button id="add_btn" class="buttons" @click="addBtnClick">Add</button>
            <button id="remove_btn" class="buttons" @click="removeItem">Remove</button>
            <button id="change_btn" class="buttons" @click="changeItem">Change</button>
        </div>
        <button id="show_btn" class="buttons" @click="updateItem">Show List</button>

        <div id="status_bar" :class="{
            unavailable: (items.length === 0),
            available: items.length
        }">
            <h1>{{ items.length }}</h1>
        </div>


        <table v-if="filteredItems.length" border="1">
            <tr >
                <td>名称</td>
                <td>位置</td>
                <td>价格</td>
                <td>所有者</td>
                <td>购入时间</td>

            </tr>
            <tr v-for="item of filteredItems">
                <td>{{ item['name'] }}</td>
                <td>{{ item['id'] }}</td>
                <td>{{ item['price'] }}</td>
                <td>{{ item['owner'] }}</td>
                <td>{{ item['time'] }}</td>
            </tr>
        </table>

        <p id="cannot_find" v-if="!filteredItems.length">找不到符合"{{ searchQuery.length >=6?searchQuery.slice(0,5)+"...":searchQuery }}"的项目</p>
        <!-- <dialog-component v-if="addItemVisible" ref="addItemDialog"></dialog-component> -->

        <br>

    </div>
</template>

<style>
#app {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    padding: 0;

}

#container {
    margin-top: 5rem;
    width: 100%;
    max-width: 500px;
    /* adjust as needed */
    padding: 0 1rem;
    box-sizing: border-box;
}

#search {
    font-size: 1.2rem;
    margin: 1rem 0;
}
#cannot_find {
    font-size: 1.7rem;
}
#status_bar {
    border-radius: 10rem;
    width: 3rem;

    box-shadow: 0rem 0rem .3rem white;

    text-align: center;

    color: azure;
    font-size: 1rem;

    text-shadow: -1px -1px 0 azure,
        1px -1px 0 azure,
        -1px 1px 0 azure,
        1px 1px 0 azure;
}

.unavailable {
    background-color: red;
}

.available {
    background-color: black;
}

#query_box {
    width: 100%;
    height: 3rem;
    font-size: 1rem;
    border-radius: 1rem;
    box-shadow: 0rem 0rem .3rem aquamarine;
    padding: 0.5rem;
}

#show_btn {
    width: 100%;
    margin-bottom: 1rem;
}

#grid {
    color: black;
    overflow-x: auto;
    font-size: 0.8rem;
}

.table {
    min-width: 100%;
}

th,
td {
    padding: 0.5rem;
}

#btns {
    display: flex;
    flex-wrap: wrap;
    justify-content: space-between;
    margin-bottom: 1rem;
}

.buttons {
    border-radius: 1rem;
    box-shadow: 0rem 0rem .7rem aquamarine;
    margin-top: 0.8rem;
    margin-bottom: 0.5rem;
    width: 48%;
    font-size: 1rem;
    height: 4rem;

}

@media (min-width: 500px) {
    #btns {
        justify-content: center;
    }

    .buttons {
        width: 30%;
    }
}
</style>