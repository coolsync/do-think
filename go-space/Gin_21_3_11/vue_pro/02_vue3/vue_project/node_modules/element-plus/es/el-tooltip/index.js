import { defineComponent, ref, h, cloneVNode } from 'vue';
import ElPopper, { defaultProps } from '../el-popper';
import { UPDATE_MODEL_EVENT } from '../utils/constants';
import throwError from '../utils/error';
import { getFirstValidNode } from '../utils/vnode';

var Tooltip = defineComponent({
    name: 'ElTooltip',
    components: {
        ElPopper,
    },
    props: Object.assign(Object.assign({}, defaultProps), { manual: {
            type: Boolean,
            default: false,
        }, modelValue: {
            type: Boolean,
            validator: (val) => {
                return typeof val === 'boolean';
            },
            default: undefined,
        }, openDelay: {
            type: Number,
            default: 0,
        }, visibleArrow: {
            type: Boolean,
            default: true,
        }, tabindex: {
            type: Number,
            default: 0,
        } }),
    emits: [UPDATE_MODEL_EVENT],
    setup(props, ctx) {
        if (props.manual && typeof props.modelValue === 'undefined') {
            throwError('[ElTooltip]', 'You need to pass a v-model to el-tooltip when `manual` is true');
        }
        const popper = ref(null);
        const onUpdateVisible = val => {
            ctx.emit(UPDATE_MODEL_EVENT, val);
        };
        const updatePopper = () => {
            return popper.value.update();
        };
        return {
            popper,
            onUpdateVisible,
            updatePopper,
        };
    },
    render() {
        const { $slots, content, manual, openDelay, onUpdateVisible, showAfter, visibleArrow, modelValue, tabindex, } = this;
        const throwErrorTip = () => {
            throwError('[ElTooltip]', 'you need to provide a valid default slot.');
        };
        const popper = h(ElPopper, Object.assign(Object.assign({}, Object.keys(defaultProps).reduce((result, key) => {
            return Object.assign(Object.assign({}, result), { [key]: this[key] });
        }, {})), { ref: 'popper', manualMode: manual, showAfter: openDelay || showAfter, showArrow: visibleArrow, visible: modelValue, 'onUpdate:visible': onUpdateVisible }), {
            default: () => ($slots.content ? $slots.content() : content),
            trigger: () => {
                if ($slots.default) {
                    const firstVnode = getFirstValidNode($slots.default(), 1);
                    if (!firstVnode)
                        throwErrorTip();
                    return cloneVNode(firstVnode, { tabindex }, true);
                }
                throwErrorTip();
            },
        });
        return popper;
    },
});

Tooltip.install = (app) => {
    app.component(Tooltip.name, Tooltip);
};
const _Tooltip = Tooltip;

export default _Tooltip;
