
<script setup>
import DemoGrid from './components/Grid.vue'
import Topbar from './components/Topbar.vue'

import axios from 'axios'

import { ref } from 'vue'


// const searchQuery = ref('')
// const gridColumns = ['name', 'id', 'price', 'owner', 'time']
// const gridData = items

const grid = ref(null)

const showGrid = ref(false)


function findItem(e) {
    alert(grid.value.filteredData[0].id);
}
function addItem() {
    alert(searchQuery[0].id);
}
function removeItem() {


}
function changeItem() {


}

function getItem() {
    const url = "http://localhost:8080/data/get"
    alert("getting message from " + url)
    axios.get(url)
        .then(res => {
            alert(res.data)
        }
        )
        .catch(error => {
            alert(error)
        })
}
</script>

<template>
    <Topbar>

    </Topbar>
    <div id="container">

        <form id="search">
            Search <input id="query_box" name="query" v-model="searchQuery">
        </form>

        <div id="btns">
            <button id="add_btn" class="buttons" @click="addItem">Add</button>
            <button id="remove_btn" class="buttons" @click="removeItem">Remove</button>
            <button id="find_btn" class="buttons" @click="findItem">Find</button>
            <button id="change_btn" class="buttons" @click="changeItem">Change</button>
        </div>
        <button id="get_btn" class="buttons" @click="getItem">Get</button>


        <!-- <DemoGrid id="grid" v-if="showGrid" :data="gridData" :columns="gridColumns" :filter-key="searchQuery" ref="grid">
        </DemoGrid> -->


        <br>

    </div>
</template>

<style>
#app {
    display: flex;

    flex-direction: column;
    justify-content: center;
    /* 水平居中 */
    align-items: center;
    padding: 0;

}

#container {
    margin-top: 5rem;
    /* adjust as needed */
}

#search {
    font-size: 1.2rem;
    margin: 1rem 0;
}

#query_box {
    width: 100%;
    height: 3rem;
    font-size: 1rem;
    border-radius: 1rem;
    box-shadow: 0rem 0rem .3rem aquamarine;
    padding: 0.5rem;
}

#get_btn {
    width: 25rem;
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
    justify-content: center;
    margin-bottom: 1rem;
}

.buttons {
    border-radius: 1rem;
    box-shadow: 0rem 0rem .7rem aquamarine;
    margin: 0.5rem;
    width: 40%;
    height: 4rem;
    font-size: 1rem;
}
</style>