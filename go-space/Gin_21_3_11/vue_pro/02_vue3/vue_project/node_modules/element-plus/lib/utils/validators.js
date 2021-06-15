'use strict';

Object.defineProperty(exports, '__esModule', { value: true });

var util = require('./util');

const isValidWidthUnit = (val) => {
    if (util.isNumber(val)) {
        return true;
    }
    else {
        return ['px', 'rem', 'em', 'vw', '%', 'vmin', 'vmax'].some(unit => val.endsWith(unit));
    }
};
const isValidComponentSize = (val) => ['', 'large', 'medium', 'small', 'mini'].includes(val);

exports.isValidComponentSize = isValidComponentSize;
exports.isValidWidthUnit = isValidWidthUnit;
