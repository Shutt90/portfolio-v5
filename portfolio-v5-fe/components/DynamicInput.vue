<template>
    <input type="{{checkInput(type)}}"/>
</template>

<script setup>
defineProps({
    type: {
        type: String,
        required: true,
    }
})

function checkInput (type) {
    if (typeof type === 'string') {
        return 'text'
    } elseif (typeof type)
}
</script>