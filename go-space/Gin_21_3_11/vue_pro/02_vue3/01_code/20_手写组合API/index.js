// 1 shallowReactive 与 reactive
const reactiveHandler = {
    get(target, prop) {
        if (prop === '_is_reactive') return true
        console.log('reactiveHandler Get: read target data');
        return Reflect.get(target, prop);
    },

    set(target, prop, value) {
        const result = Reflect.set(target, prop, value)
        console.log('set method, 数据已更新, 去更新界面')
        return result
    },

    deleteProperty(target, prop) {
        const result = Reflect.deleteProperty(target, prop)
        console.log('deleteProperty method, 数据已删除, 去更新界面')
        return result
    },
}

// 自定义shallowReactive
function shallowReactive(target) {
    if (target && typeof target === 'object') {
        return new Proxy(target, reactiveHandler)
    }
    return target
}

// 自定义reactive
function reactive(target) {
    if (target && typeof target === 'object') {
        if (target instanceof Array) { // 数组
            target.forEach((item, index) => {
                target[index] = reactive(item)
            })
        } else { // 对象
            Object.keys(target).forEach(key => {
                target[key] = reactive(target[key])
            })
        }

        const proxy = new Proxy(target, reactiveHandler)
        return proxy
    }

    return target
}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

//2 shallowReadonly 与 readonly
const readonlyHandler = {
    get(target, prop) {
        if (prop === '_is_readonly') return true
        console.log('readonlyHandler get: read data', target, prop)
        return Reflect.get(target, prop)
    },

    set() {
        console.warn('set: 只读的, 不能修改')
        return true
    },

    deleteProperty() {
        console.warn('deleteProperty: 只读的, 不能删除')
        return true
    },
}

// 自定义shallowReadonly
function shallowReadonly(target_obj) {
    if (target_obj && typeof target_obj === 'object') {
        return new Proxy(target_obj, readonlyHandler)
    }
    return target_obj
}

// 自定义readonly
function readonly(target) {
    if (target && typeof target === 'object') {
        if (Array.isArray(target)) { // 数组
            target.forEach((item, index) => {
                target[index] = readonly(item)
            })
        } else { // 对象
            Object.keys(target).forEach(key => {
                target[key] = readonly(target[key])
            })
        }

        const proxy = new Proxy(target, readonlyHandler)
        return proxy
    }

    return target
}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
// 3 shallowRef 与 ref
// 自定义shallowRef
function shallowRef(target) {
    const result = {
        _value: target, // 用来保存数据的内部属性
        _is_ref: true, // 用来标识是ref对象
        get value() {
            console.log('Get value: data')
            return this._value
        },
        set value(val) {
            console.log('Set value 数据已更新, 去更新界面')
            this._value = val
        }
    }
    return result
}

// 自定义ref
function ref(target) {
    // if (target && typeof target === 'object') {
    target = reactive(target)
    // }
    return {
        _value: target, // 用来保存数据的内部属性
        _is_ref: true, // 用来标识是ref对象
        get value() {
            console.log('Get value: data')
            return this._value
        },
        set value(val) {
            this._value = val
            console.log('Set value 数据已更新, 去更新界面')
        }
    }
    // return result
}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
// 判断是否是ref对象
function isRef(obj) {
    return obj && obj._is_ref
}

// 判断是否是reactive对象
function isReactive(obj) {
    return obj && obj._is_reactive
}
// 判断是否是readonly对象
function isReadonly(obj) {
    return obj && obj._is_readonly
}

// 是否是reactive或readonly产生的代理对象
function isProxy(obj) {
    return isReactive(obj) || isReadonly(obj)
}