import axios from "axios";
import { ref } from "vue";
export default function <T>(url: string) {
    const loading = ref(true);
    const errMsg = ref('');
    const ret = ref<T | null>(null)

    axios.get(url).then(res => {
        loading.value = false
        ret.value = res.data
    })
    .catch(err => {
        loading.value = false;
        errMsg.value = err.message || 'unknown err';
    })

    return {
        loading,
        errMsg,
        ret,
    }
}