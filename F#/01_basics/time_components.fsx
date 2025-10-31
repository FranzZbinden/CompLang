(*

 * Author: Franz Zbinden
 * Course: COTI 4039-LH1
 * Date: 09/08/2025
 * Purpose: This program computes and displays the components (hours, minutes,
 *          and seconds) of a moment in time.
 *)



let timeComponents elapsedSecs =    
    let secsPerHour = 3600

    let hrs = elapsedSecs / secsPerHour
    let remainingSecs = elapsedSecs % secsPerHour

    let secsPerMinute = 60
    let mins = remainingSecs / secsPerMinute
    let secs = remainingSecs % secsPerMinute
    hrs, mins, secs

let elapsedSeconds = 12345
printfn "The number of elapsed seconds is %d" elapsedSeconds
let hours, minutes, seconds = timeComponents elapsedSeconds
printfn "\nThis is %d hours, %d minutes, and %d seconds" hours minutes seconds

