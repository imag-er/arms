import {lab_places} from './labPlaces.js'
import {classes} from './classes.js'

export function makePlaceString(item_id:Number) {
    // place number str
    var pns = item_id.toString()
    
    var res = `${lab_places[pns[0]].label}${lab_places[pns[0]].children[pns[1]].label}${pns[2]}行${pns[3]}列`
    return res
}


export function makeClassString(item_id) {
    var a = Math.floor(item_id / 1000) % 10
    var b = Math.floor(item_id / 100) % 10
     

    var res = `${classes[a].label} - ${classes[a].children[b].label}`
    return res
}


export function makeDateString() {
    var today = new Date();
    var year = today.getFullYear();
    var month = today.getMonth() + 1;
    var day = today.getDate();

    //将月份和日期补零
    if (month < 10) {
        month = '0' + month;
    }
    if (day < 10) {
        day = '0' + day;
    }

    var formattedDate = '' + year + month + day;

    return formattedDate
}