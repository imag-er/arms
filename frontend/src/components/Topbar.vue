<script setup>


</script>

<template>
    <div class="topbar">
        <h1>
            AILab资产管理系统
        </h1>


        <a>登录</a>|<a>注册</a>

    </div>
</template>

<style>
.topbar {
    background-color: rgba(144, 237, 219, 0.664);
    width: 100%;
    display: flex;
    align-items: center;
    justify-content: space-around;
    position: fixed;
    top: 0;
    left: 0;
    z-index: 9999;
    box-shadow: 0 0.1rem 25rem rgba(255, 255, 255, 0.3), 0 0 20px rgba(255, 255, 255, 0.5);

    text-shadow: -1px -1px 0 #000,
        1px -1px 0 #000,
        -1px 1px 0 #000,
        1px 1px 0 #000;

}

.topbar::before {
    content: '';
    width: 0.3rem;
    height: 100%;
    background-color: aquamarine;
    position: absolute;
    left: 0;
}

.topbar>h1 {
    color: aquamarine;
    font-size: 1.3rem;
    ;
}

.topbar>a {
    color: aqua;
    font-size: 1rem;
}

.topbar>* {
    font-size: 1.5rem;
    margin: 0.1rem;
}
</style>