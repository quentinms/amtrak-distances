'use strict'

function getDistance(id, e) {
    // TODO check undefined
    const start = document.getElementById(`select-from-${id}`).value
    const end = document.getElementById(`select-to-${id}`).value
    const dist = Math.abs(end-start)
    document.getElementById(`dist-${id}`).innerText = `Distance: ${dist} miles (${Math.round(dist * 1.609344)} km)`
}
