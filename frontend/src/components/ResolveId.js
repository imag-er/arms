
export function makePlaceString(item_id) {
    // place number str
    var pns = item_id.toString()
    var res = `${pns[0]}排${pns[1]}区${pns[2]}行${pns[3]}列`
    return res
}


export function makeClassString(item_id) {
    
}
