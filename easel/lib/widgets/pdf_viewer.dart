import 'dart:async';
import 'dart:io';

import 'package:easel_flutter/easel_provider.dart';
import 'package:easel_flutter/screens/clippers/right_triangle_clipper.dart'
    as clipper;
import 'package:easel_flutter/screens/clippers/right_triangle_clipper.dart';
import 'package:easel_flutter/screens/clippers/small_bottom_corner_clipper.dart';
import 'package:easel_flutter/utils/constants.dart';
import 'package:easel_flutter/utils/easel_app_theme.dart';
import 'package:easel_flutter/utils/extension_util.dart';
import 'package:easel_flutter/utils/route_util.dart';
import 'package:easel_flutter/widgets/pdf_viewer_full_half_screen.dart';
import 'package:easy_localization/easy_localization.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flutter_pdfview/flutter_pdfview.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:flutter_svg/svg.dart';
import 'package:get_it/get_it.dart';
import 'package:path_provider/path_provider.dart';
import 'package:provider/provider.dart';

import '../generated/locale_keys.g.dart';

class PdfViewer extends StatefulWidget {
  final File? file;
  final String? fileUrl;
  final bool previewFlag;

  const PdfViewer(
      {Key? key, this.file, required this.previewFlag, this.fileUrl})
      : super(key: key);

  @override
  State<PdfViewer> createState() => _PdfViewerState();
}

class _PdfViewerState extends State<PdfViewer> with WidgetsBindingObserver {
  EaselProvider get easelProvider => GetIt.I.get();

  late String doc;
  bool _isLoading = true;

  final Completer<PDFViewController> _controller =
  Completer<PDFViewController>();
  int? pages = 0;
  bool isReady = false;

  String errorMsg = "";
  @override
  void initState() {
    scheduleMicrotask(() {
      easelProvider.setPdfThumbnail(null);
    });
    initializeDoc();
    super.initState();
  }

  Future initializeDoc() async {
    if (widget.file == null && widget.fileUrl == null) {
      errorMsg = LocaleKeys.no_pdf_file.tr();
      _isLoading = false;
      setState(() {});
      return;
    }
    if (widget.file != null) {
      doc = widget.file!.path;
      _isLoading = false;
      setState(() {});
    } else {
      final Completer<File> completer = Completer();
      try {
          final url = widget.fileUrl!;
          final filename = url.substring(url.lastIndexOf("/") + 1);
          final request = await HttpClient().getUrl(Uri.parse(url));
          final response = await request.close();
          final bytes = await consolidateHttpClientResponseBytes(response);
          final dir = await getApplicationDocumentsDirectory();
          final File file = File("${dir.path}/$filename");

          await file.writeAsBytes(bytes, flush: true);
          completer.complete(file);
        } catch (e) {
          throw Exception('Error parsing asset file!');
        }

        final File file = await completer.future;
        doc = file.path;
        _isLoading = false;
        setState(() {});
    }

  }

  @override
  Widget build(BuildContext context) {
    return ChangeNotifierProvider<EaselProvider>.value(
      value: easelProvider,
      child: Center(
        child: PdfViewerFullOrHalf(
          pdfViewerFullScreen: (context) {
            return errorMsg.isNotEmpty
                ? Text(
                    errorMsg,
                    style: const TextStyle(color: EaselAppTheme.kWhite),
                  )
                : Padding(
                    padding: EdgeInsets.only(top: 100.h, bottom: 145.h),
                    child: PDFView(
                      filePath: doc,
                      swipeHorizontal: true,
                      onRender: (pgs) {
                        setState(() {
                          pages = pgs;
                          isReady = true;
                        });
                      },
                      onViewCreated: (PDFViewController pdfViewController) {
                        _controller.complete(pdfViewController);
                      },
                    ),
                  );
          },
          pdfViewerHalfScreen: (context) {
            return SingleChildScrollView(
                child: Column(
              children: [
                Padding(
                  padding: EdgeInsets.symmetric(horizontal: 4.w),
                  child: SizedBox(
                      height: 200.h,
                      child: errorMsg.isNotEmpty
                          ? Center(
                              child: Text(
                                errorMsg,
                              ),
                            )
                          : Stack(
                              children: [
                                PDFView(
                                  filePath: doc,
                                  swipeHorizontal: true,
                                  onRender: (pgs) {
                                    setState(() {
                                      pages = pgs;
                                      isReady = true;
                                    });
                                  },
                                  onViewCreated: (PDFViewController pdfViewController) {
                                    _controller.complete(pdfViewController);
                                  },
                                ),

                                _buildPdfFullScreenIcon()
                              ],
                            )),
                ),
                SizedBox(
                  height: 50.h,
                ),
                _buildThumbnailButton(),
              ],
            ));
          },
          previewFlag: widget.previewFlag,
          isLoading: _isLoading,
        ),
      ),
    );
  }

  Widget _buildPdfFullScreenIcon() {
    return Positioned(
      left: 5,
      bottom: 0,
      child: ClipPath(
        clipper: RightTriangleClipper(
            orientation: clipper.Orientation.orientationNE),
        child: InkWell(
          onTap: () {
            Navigator.pushNamed(context, RouteUtil.kPdfFullScreen,
                arguments: [doc]);
          },
          child: Container(
            width: 30.w,
            height: 30.w,
            alignment: Alignment.bottomLeft,
            color: EaselAppTheme.kLightRed,
            child: Padding(
              padding: EdgeInsets.all(5.w),
              child: RotationTransition(
                turns: const AlwaysStoppedAnimation(0.25),
                child: SvgPicture.asset(
                  PngUtils.kFullScreenIcon,
                  fit: BoxFit.fill,
                  width: 8.w,
                  height: 8.w,
                  alignment: Alignment.bottomLeft,
                ),
              ),
            ),
          ),
        ),
      ),
    );
  }

  Widget _buildThumbnailButton() {
    return Align(
      alignment: Alignment.bottomLeft,
      child: Padding(
        padding: EdgeInsets.only(left: 20.h),
        child: SizedBox(
          height: 120.h,
          width: 120.w,
          child: InkWell(
            onTap: () {
              if (errorMsg.isNotEmpty) {
                  LocaleKeys.first_pick_pdf.tr().show();
                return;
              }
              easelProvider.onPdfThumbnailPicked();
            },
            child: easelProvider.pdfThumbnail != null
                ? ClipPath(
                    clipper: RightSmallBottomClipper(),
                    child: Container(
                        height: 60.h,
                        width: 60.w,
                        margin: EdgeInsets.only(left: 10.w),
                        child: Image.file(
                          easelProvider.pdfThumbnail!,
                          height: 60.h,
                          width: 60.w,
                          fit: BoxFit.cover,
                        )),
                  )
                : SvgPicture.asset(PngUtils.kUploadThumbnail),
          ),
        ),
      ),
    );
  }

  bool shouldShowThumbnailButton() {
    return !widget.previewFlag;
  }
}
