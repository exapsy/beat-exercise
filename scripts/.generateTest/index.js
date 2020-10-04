const fs = require('fs')

const ERR_USAGE = 1
const ERR_OPEN_FILE = 2
const ERR_WRITE_FILE = 3

const { argv } = process
if (argv.length < 3) {
    console.error('Usage: <lines> <rides> <output>')
    process.exit(ERR_USAGE)
}
let [,, lines, rides, outputFile] = argv

fs.open(outputFile, 'w', (err, fd) => {
    if (err) {
        console.error(err)
        process.exit(ERR_OPEN_FILE)
    }
    const linesEachRide = lines / rides
    for (let ride = 0; ride < rides; ride++) {
        const rideOutput = generateRide(ride, linesEachRide)
        fs.writeFile(fd, rideOutput, (err) => {
            if (err) {
                console.error(err)
                process.exit(ERR_WRITE_FILE)
            }
        })
    }
})

const generateRide = (id, totalSegments) => {
    let rideOutput = ''
    let _date = new Date()
    let originalLatitude = (Math.random() * 100) % 90
    let originalLongitude = (Math.random() * 200) % 180
    for (let i = 0; i < totalSegments; i++) {
        _date.setTime(_date.getTime() + Math.random() * 10)
        const date = Date.parse(_date.toISOString())
        const latitude = (originalLatitude) + (1 * Math.random())
        const longitude = (originalLongitude) + (1 * Math.random())
        rideOutput = rideOutput.concat(`${id},${latitude},${longitude},${date}\r\n`)
    }
    return rideOutput
}