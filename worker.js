self.onmessage = async function (event) {
    let clampedArray = new Uint8ClampedArray(await event.data.arrayBuffer());
    let imgData = new ImageData(clampedArray, 1920, 1080);
    let bitmap = await createImageBitmap(imgData);
    self.postMessage(bitmap);
}
