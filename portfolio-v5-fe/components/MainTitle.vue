<template>
    <h1 class="title">{{ title }}</h1>
</template>

<script setup>
// import gsap from 'gsap';

// gsap.from(".title", {
//     opacity: 0,
// })
defineProps({
    title: {
        required: true,
        type: String,
    }
})
</script>