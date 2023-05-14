<script setup>
const addItemVisible = ref(false)


function showDialog(data) {
    addItemVisible = true;
    console .log(data)
}


function closeDialog(data) {
    addItemVisible = false;
    console .log(data)
}

defineExpose(
    [
        showDialog
    ]
)
</script>

<template>
    <el-dialog title="添加物品" :visible.sync="addItemVisible" width="35%">
        测试
    </el-dialog>
</template>

<style></style>