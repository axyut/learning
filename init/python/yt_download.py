from pytube import YouTube 

link = input(" Enter Video Url: ")
# reso = input("Enter required Quality: ")
yt= YouTube(link)

print(" Title: ", yt.title)
print(" Views: ", yt.views)
print(" Downloading...")

video = yt.streams.get_highest_resolution()
video.download()
input(f" {video} is downloaded! (Highest Available Quality Video) \n Press Enter to exit.")