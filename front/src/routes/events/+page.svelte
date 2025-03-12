<script>
    import {events, eventsError, getEvents} from "$lib/stores/events.js";
    import {onMount} from "svelte";
    import {SvelteMap} from "svelte/reactivity";
    import {getResource} from "$lib/stores/resources.js";
    import Calendar from '@event-calendar/core';
    import DayGrid from '@event-calendar/day-grid';


    const pickColorsFrom = [
        '#F6DC00',
        '#FFFFFF',
        '#AB69AD',
        '#FAC000',
        '#E7A2FD',
        '#DE9800',
    ]
    let resourceColors = new SvelteMap()
    let eventsResources = new SvelteMap()
    let plugins = [DayGrid];
    let options = {
        view: 'dayGridMonth',
        events: [],
        editable: false,
        eventContent: (info) => {
            let titles = info.event.title.split('|')
            let desc = titles[1].replace(/^[\\n]+/g,'').replace(/\\n/g,'&#013;')
            return {html: `<div style="${getEventColors(info.event)}"><span class="fs-5 me-2 fw-bold">${info.event.start.getHours()}:${info.event.start.getMinutes()}</span><span class="fs-5" title="${desc}">${titles[0]} <br/> <i class="fs-6">${titles[2]}</i></span></div>`}
        },
        eventBackgroundColor : '#593196'
    };
    // maybe see
    //   eventClick
    //   eventMouseEnter / eventMouseLeave for hovering event
    //   loading
    //   pointer : enables mouse pointer
    //   Methodsto : invoke calendar methods
    //   eventBackgroundColor can only be string :/


    onMount(() => {
        getEvents()
    })

    function getEventColors(event){
        let colors = []
        let color = ""
        event.resourceIds.forEach((r) => {
            if(!resourceColors.get(r)){
                // pick new color
                resourceColors.set(r, pickColorsFrom[resourceColors.size])
            }
            colors.push(resourceColors.get(r))
        })
        if(colors.length > 1){
            let percent = 100 / colors.length
            // background-image: linear-gradient(red 50%, blue 50%);
            color = "color: transparent;background-clip: text;background-image: linear-gradient(to right,"
            colors.forEach((c) => {
                color += `${c} ${percent}%,`
            })
            color = color.substring(0, color.length - 1); // remove last ','
            color += ");"
        } else {
            color = `color: ${colors[0]};`
        }
        return color
    }

    events.subscribe((values) => {
        let list = []
        values.forEach((value) => {
            // Set event for calendar component
            let e = {
                id: value.id,
                //allDay: false,
                start: new Date(value.start),
                end: new Date(value.end),
                title: `${value.name}|${value.description}|${value.location}`, //todo see with desc
                //editable: false,
                //startEditable: false,
                //durationEditable: false,
                resourceIds: value.resourceIds, //todo see ??
                //display: 'auto',
                //backgroundColor: '', // see also eventBackgroundColor
                //textColor: '', // see also eventTextColor
                //color: '', // alias for backgrounColor
                //classNames: [], // see also eventClassNames
                //styles: [],
                //extendedProps: {},
            }
            list.push(e)

            // Get resources linked to event
            // really not optimized
            value.resourceIds.forEach((rId) => {
                if(!eventsResources.get(rId)){
                    getResource(rId)
                        .then((data) => {
                            eventsResources.set(rId, data)
                        })
                        .catch(()=>{})
                }
            })
        })
        options.events = list;
    })
</script>

<div class="d-flex flex-column">
    <h1 class="text-center color-yellow mb-lg-5">Events</h1>
    {#if $eventsError}
        <div  class="text-center text-danger form-text fs-5">
            An error occurred : {$eventsError}
        </div>
    {/if}

    <div class="d-flex mb-4 justify-content-center">
        {#each resourceColors as [r, c]}
            <div class="d-flex me-3">
                <span class="p-3 me-2" style="{`background-color: ${c};`}"></span>
                <span>{eventsResources.get(r) ? eventsResources.get(r).name : r}</span>
            </div>
        {/each}
    </div>

    <Calendar {plugins} {options}/>
</div>