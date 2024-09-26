# Google Photos Downloader for rclone

This is a Google Photos downloader for use with rclone.

The Google Photos API delivers images and video which aren't full resolution, and/or have EXIF data missing (see [#112096115](https://issuetracker.google.com/issues/112096115) and [#113672044](https://issuetracker.google.com/issues/113672044))

However if you use this proxy then you can download original, unchanged images as uploaded by you.

This runs a headless browser in the background with an HTTP server which [rclone](https://rclone.org) which uses the Google Photos website to fetch the original resolution images.

## Usage

First [install rclone](https://rclone.org/install/) and set it up with [google photos](https://rclone.org/googlephotos/).

You will need rclone version v1.69 or later. If v1.69 hasn't been release yet then please use [the latest beta](https://beta.rclone.org/).

Next download the latest gphotosdl binary from the [releases page](https://github.com/rclone/gphotosdl/releases/latest).

You will need to run like this first. This will open a browser window which you should use to login to google photos - then close the browser window. You may have to do this again if the integration stops working.

    gphotosdl -login

Once you have done this you can run this to run the proxy.

    gphotosdl

Then supply the parameter `--gphotos-proxy "http://localhost:8282"` to make rclone use the proxy. For example

    rclone copy -vvP --gphotos-proxy "http://localhost:8282" gphotos:media/by-month/2024/2024-09/ /tmp/high-res-media/

Run the `gphotosdl` command with the `-debug` flag for more info and the `-show` flag to see the browser that it is using. These are essential if you are trying to debug a problem.

    gphotosdl -debug -show

## Troubleshooting

You can't run more than one proxy at once. If you get the error 

    browser launch: [launcher] Failed to get the debug url: Opening in existing browser session.

Then there is another `gphotosdl` running or there is an orphan browser process you will have to kill.

## Limitations

- Currently only fetches one image at once. Conceivably could make multiple tabs in the browser to fetch more than one at once.
- More error checking needed - if it goes wrong then it will hang forever most likely
- Currently the browser only has one profile so this can only be used with one google photos user. This is easy to fix.

## License

This is free software under the terms of the MIT license (check the LICENSE file included in this package).

## Contact and support

The project website is at:

- https://github.com/rclone/gphotosdl

There you can file bug reports, ask for help or contribute patches.
