import ConnectHub from '../SpindlyHubs.js'

const hub_name = 'GlobalHub';

export function GlobalHub(hub_instance_id) {
    let SpindlyStore = ConnectHub(hub_name, hub_instance_id);
    return {
        AppName: SpindlyStore("AppName"),
        Msg: SpindlyStore("Msg"),
        SrcInterface: SpindlyStore("SrcInterface"),
        DstInterface: SpindlyStore("DstInterface"),
        Pass: SpindlyStore("Pass"),
        Name: SpindlyStore("Name"),
        NetworkInterfaces: SpindlyStore("NetworkInterfaces"),
        Refresh: SpindlyStore("Refresh"),
        Start: SpindlyStore("Start"),
        IsRunning: SpindlyStore("IsRunning"),
    }
}

export let Global = GlobalHub("Global");
