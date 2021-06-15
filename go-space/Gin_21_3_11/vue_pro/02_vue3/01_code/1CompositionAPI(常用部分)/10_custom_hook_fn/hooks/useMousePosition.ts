import { onMounted, onUnmounted, ref } from "vue";

export default function () {
    const x = ref(-1);
    const y = ref(-1);

    const clickHandler = (e: MouseEvent) => {
        x.value = e.pageX;
        y.value = e.pageY;
    }

    onMounted(() => {
        addEventListener("click", clickHandler)
    });

    onUnmounted(() => {
        removeEventListener("click", clickHandler)
    })
    return {
        x,
        y
    }
}