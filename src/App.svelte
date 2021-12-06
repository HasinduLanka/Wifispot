<script>
  import { Global } from "./stores/global.spindlyhubs.js";

  let {
    AppName,
    Msg,
    NetworkInterfaces,
    Refresh,
    SrcInterface,
    DstInterface,
    Start,
    IsRunning,
    Name,
    Pass,
  } = Global;

  let failed = false;
  if (!$NetworkInterfaces || !$NetworkInterfaces.length) {
    $NetworkInterfaces = [];
    failed = true;
    $Msg = "No network interfaces found.";
  }

  function RefreshClick() {
    $Refresh = !$Refresh;
  }

  RefreshClick();

  function StartClick() {
    $Start = !$Start;
  }

  // Github copilot is a god.

  if (!$Name) {
    $Name = localStorage.getItem("hotspot-name");
  }

  if (!$Pass) {
    $Pass = localStorage.getItem("hotspot-pass");
  }

  if (!$SrcInterface) {
    $SrcInterface = localStorage.getItem("hotspot-src-interface");
  }

  if (!$DstInterface) {
    $DstInterface = localStorage.getItem("hotspot-dst-interface");
  }

  Name.subscribe((val) => localStorage.setItem("hotspot-name", val));
  Pass.subscribe((val) => localStorage.setItem("hotspot-pass", val));

  SrcInterface.subscribe((val) =>
    localStorage.setItem("hotspot-src-interface", val)
  );

  DstInterface.subscribe((val) =>
    localStorage.setItem("hotspot-dst-interface", val)
  );
</script>

<h1>{$AppName}</h1>

<p>{$Msg}</p>

<div>
  <input type="button" value="Refresh" on:click={RefreshClick} />
</div>
<br />

<div>
  Share internet from :
  <select id="network-interface-selector" bind:value={$SrcInterface}>
    {#each $NetworkInterfaces as NI}
      <option value={NI}>{NI}</option>
    {/each}
  </select>
</div>

<div>
  Share internet to :
  <select id="network-interface-selector" bind:value={$DstInterface}>
    {#each $NetworkInterfaces as NI}
      <option value={NI}>{NI}</option>
    {/each}
  </select>
</div>

<div>
  <input type="text" bind:value={$Name} placeholder="WiFi Hotspot name" />
</div>

<div>
  <input type="text" bind:value={$Pass} placeholder="Hotspot password" />
</div>

<h3>{$IsRunning ? "Running" : ""}</h3>

<input
  type="button"
  value={$IsRunning ? "Stop" : "Start"}
  on:click={StartClick}
/>

<br />
