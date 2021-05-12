'use strict';

Object.defineProperty(exports, '__esModule', { value: true });

var vue = require('vue');
var ElPopper = require('../el-popper');
var constants = require('../utils/constants');
var throwError = require('../utils/error');
var vnode = require('../utils/vnode');

function _interopDefaultLegacy (e) { return e && typeof e === 'object' && 'default' in e ? e : { 'default': e }; }

var ElPopper__default = /*#__PURE__*/_interopDefaultLegacy(ElPopper);
var throwError__default = /*#__PURE__*/_interopDefaultLegacy(throwError);

var Tooltip = vue.defineComponent({
    name: 'ElTooltip',
    components: {
        ElPopper: ElPopper__default['default'],
    },
    props: Object.assign(Object.assign({}, ElPopper.defaultProps), { manual: {
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
    emits: [constants.UPDATE_MODEL_EVENT],
    setup(props, ctx) {
        if (props.manual && typeof props.modelValue === 'undefined') {
            throwError__default['default']('[ElTooltip]', 'You need to pass a v-model to el-tooltip when `manual` is true');
        }
        const popper = vue.ref(null);
        const onUpdateVisible = val => {
            ctx.emit(constants.UPDATE_MODEL_EVENT, val);
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
            throwError__default['default']('[ElTooltip]', 'you need to provide a valid default slot.');
        };
        const popper = vue.h(ElPopper__default['default'], Object.assign(Object.assign({}, Object.keys(ElPopper.defaultProps).reduce((result, key) => {
            return Object.assign(Object.assign({}, result), { [key]: this[key] });
        }, {})), { ref: 'popper', manualMode: manual, showAfter: openDelay || showAfter, showArrow: visibleArrow, visible: modelValue, 'onUpdate:visible': onUpdateVisible }), {
            default: () => ($slots.content ? $slots.content() : content),
            trigger: () => {
                if ($slots.default) {
                    const firstVnode = vnode.getFirstValidNode($slots.default(), 1);
                    if (!firstVnode)
                        throwErrorTip();
                    return vue.cloneVNode(firstVnode, { tabindex }, true);
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

exports.default = _Tooltip;
