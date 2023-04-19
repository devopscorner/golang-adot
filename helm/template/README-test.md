## README How To Test

### Prerequirements

- Add Helm repository

  ```
  ### LAB ###
  helm s3 init s3://devopscorner-adot-chart/lab
  AWS_REGION=us-west-2 helm repo add devopscorner-adot-lab s3://devopscorner-adot-chart/lab

  ### STAGING ###
  helm s3 init s3://devopscorner-adot-chart/staging
  AWS_REGION=us-west-2 helm repo add devopscorner-adot-staging s3://devopscorner-adot-chart/staging

  ### PRODUCTION ###
  helm s3 init s3://devopscorner-adot-chart/prod
  AWS_REGION=us-west-2 helm repo add devopscorner-adot s3://devopscorner-adot-chart/prod

  helm repo update
  ```

### Helm Production

- Test Rendering

  ```
  cd prod
  helm template api -f api/values.yaml
  helm template backend -f backend/values.yaml
  helm template frontend -f frontend/values.yaml
  helm template stateful -f stateful/values.yaml
  helm template config -f config/values.yaml
  helm template secret -f secret/values.yaml
  ```

- Test Deploy

  ```
  helmfile -f test/prod/prod-helm-120-template.yml
  ```

### Helm Staging and Helm Lab

- Test Rendering

  ```
  cd lab
  helm template api -f api/values.yaml
  helm template backend -f backend/values.yaml
  helm template frontend -f frontend/values.yaml
  helm template stateful -f stateful/values.yaml
  helm template config -f config/values.yaml
  helm template secret -f secret/values.yaml

  cd staging
  helm template api -f api/values.yaml
  helm template backend -f backend/values.yaml
  helm template frontend -f frontend/values.yaml
  helm template stateful -f stateful/values.yaml
  helm template config -f config/values.yaml
  helm template secret -f secret/values.yaml
  ```

- Test Deploy

  ```
  helmfile -f test/staging-lab/lab-helm-120-template.yml
  helmfile -f test/staging-lab/staging-helm-120-template.yml
  ```
