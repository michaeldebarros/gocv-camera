FROM golang:1.11.1-stretch as base
WORKDIR /camera
COPY . .
ENV GO_CXXFLAGS=-std=c++11
ENV CGO_CPPFLAGS=-I/usr/local/Cellar/opencv/3.4.3/include
ENV CGO_CPPFLAGS="-L/usr/local/Cellar/opencv/3.4.3/lib -lopencv_stitching -lopencv_superres -lopencv_videostab -lopencv_aruco -lopencv_bgsegm -lopencv_bioinspired -lopencv_ccalib -lopencv_dnn_objdetect -lopencv_dpm -lopencv_face -lopencv_photo -lopencv_fuzzy -lopencv_hfs -lopencv_img_hash -lopencv_line_descriptor -lopencv_optflow -lopencv_reg -lopencv_rgbd -lopencv_saliency -lopencv_stereo -lopencv_structured_light -lopencv_phase_unwrapping -lopencv_surface_matching -lopencv_tracking -lopencv_datasets -lopencv_dnn -lopencv_plot -lopencv_xfeatures2d -lopencv_shape -lopencv_video -lopencv_ml -lopencv_ximgproc -lopencv_calib3d -lopencv_features2d -lopencv_highgui -lopencv_videoio -lopencv_flann -lopencv_xobjdetect -lopencv_imgcodecs -lopencv_objdetect -lopencv_xphoto -lopencv_imgproc -lopencv_core"
RUN go get
RUN go build -tags customenv -o main .

FROM scratch
COPY --from=base /camera/main /camera

EXPOSE 8080

CMD ["/camera"]

