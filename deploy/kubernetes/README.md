# Kubernetes

Run TexWeave as a **Job**: input from ConfigMap, output to emptyDir (copy out with `kubectl cp`).

1. Copy `secret.example.yaml` to `secret.yaml`, set your API key(s), then:
   ```bash
   kubectl apply -f deploy/kubernetes/secret.yaml
   ```
2. Apply ConfigMap and Job:
   ```bash
   kubectl apply -f deploy/kubernetes/configmap-input.yaml -f deploy/kubernetes/job.yaml
   ```
3. Wait and copy output:
   ```bash
   kubectl wait --for=condition=complete job/texweave-generate --timeout=120s
   kubectl cp $(kubectl get pods -l job-name=texweave-generate -o jsonpath='{.items[0].metadata.name}'):/output/output.tex output.tex
   ```

Image: `ghcr.io/ali-m07/texweave:latest` (built by GitHub Actions). Do not commit `secret.yaml`.
